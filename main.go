package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/matyix/echo/conf"
	"github.com/matyix/echo/log"
	"github.com/matyix/echo/middlewares"
	"github.com/olahol/melody"
	"net/http"
	"time"
)

var GitCommit string
const Version = "0.1.0"
var VersionPrerelease = "dev"

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		return
	}

	log.WithFields(log.Fields{
		"EventName": "get_config_env_vars",
	}).Debug("Reading configuration from env vars")
	cfg := conf.Config()

	log.WithFields(log.Fields{
		"EventName": "set_gin_mode",
		"Mode":      cfg.GetString("mode"),
	}).Debug("Setting gin_ mode to ", cfg.GetString("mode"))
	gin.SetMode(cfg.GetString("mode"))

	ginweb := gin.New()
	melodyws:= melody.New()

	//Web Sockets
	authorized.GET("/channel/:name/websockets", func(c *gin.Context) {
		melodyws.HandleRequest(c.Writer, c.Request)
	})

	ginweb.Use(ginrus.Ginrus(log.NewLogger(cfg), time.RFC3339, true))
	ginweb.Use(gin.Recovery())

	melodyws.HandleConnect(func(session *melody.Session) {
		log.WithFields(log.Fields{
			"EventName":     "websockets_client_connect",
			"RemoteAddress": session.Request.RemoteAddr,
		}).Debug("new websockets client connected ", session.Request.RemoteAddr)
	})

	melodyws.HandleDisconnect(func(session *melody.Session) {
		log.WithFields(log.Fields{
			"EventName":     "websockets_client_disconnect",
			"RemoteAddress": session.Request.RemoteAddr,
		}).Debug("WebSockets client disconnected ", session.Request.RemoteAddr)
	})

	melodyws.HandleError(func(session *melody.Session, err error) {
		log.WithFields(log.Fields{
			"EventName":     "websockets_error",
			"RemoteAddress": session.Request.RemoteAddr,
			"Error":         err.Error(),
		}).Error("error ocurred with WebSockets client ", err.Error())
	})

	melodyws.HandleMessage(func(session *melody.Session, msg []byte) {
		melodyws.BroadcastFilter(msg, func(q *melody.Session) bool {
			msgStr := string(msg[:])

			log.WithFields(log.Fields{
				"EventName": "websockets_client_message",
				"Message":   msgStr,
				"Channel":   q.Request.URL.Path,
			}).Debug("reseived message: ", msgStr)

			return q.Request.URL.Path == session.Request.URL.Path
		})
	})

	// Start Server
	fmt.Printf("STARTED", "started" )
	log.WithFields(log.Fields{
		"EventName":         "start",
		"ListenAddress":     cfg.GetString("listen_address"),
		"GitCommit":         GitCommit,
		"Version":           Version,
		"VersionPrerelease": VersionPrerelease,
		//"TLS": cfg.GetBool("secure"),
		"ReadTimeout": cfg.GetDuration("read_timeout"),
		"WriteTimeout": cfg.GetDuration("write_timeout"),
		"MaxHeaderBytes": cfg.GetInt("max_header_bytes"),
	}).Info("starting server and listening on ", cfg.GetString("listen_address"))

	server := &http.Server{
		Addr:           cfg.GetString("listen_address"),
		Handler:        ginweb,
		ReadTimeout:    cfg.GetDuration("read_timeout"),
		WriteTimeout:   cfg.GetDuration("write_timeout"),
		MaxHeaderBytes: cfg.GetInt("max_header_bytes"),
		//TLSConfig:      tlscfg,

	}

	switch cfg.GetBool("secure") {
	case true:
		// Key or DH parameter strength >= 4096 bits (e.g., 4096)
		log.Fatal(server.ListenAndServeTLS(cfg.GetString("cert_file"), cfg.GetString("key_file")))
	default:
		log.Fatal(server.ListenAndServe())
	}
}

