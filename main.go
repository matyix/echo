package main

import (
	"flag"
	"fmt"
	"github.com/matyix/echo/conf"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/gin-gonic/contrib/ginrus"
	"time"
	"github.com/matyix/echo/middlewares"
	"crypto/tls"
	"net/http"
	"github.com/matyix/echo/log"
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
	}).Debug("Setting gin mode to ", cfg.GetString("mode"))
	gin.SetMode(cfg.GetString("mode"))

	//r := gin.New()
	//m := melody.New()

}
