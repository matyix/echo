package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matyix/echo/conf"
	"github.com/matyix/echo/log"
)


func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := stripBearerPrefixFromTokenString(c.Request.Header.Get("Authorization"))

		if err != nil {

			log.WithFields(log.Fields{
				"EventName": "middleware_auth_error",
				"Error":     err.Error(),
			}).Error("error while stripping bearer prefix ", err.Error())

			c.AbortWithError(401, err)
		}

		cfg := config.Config()
		if token != cfg.GetString("secret") {

			log.WithFields(log.Fields{
				"EventName": "middleware_auth_denied",
				"Token":     token,
			}).Error("Invalid token ", token)

			c.AbortWithStatus(401)
		}
	}
}
