package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/api-gateway/internal/pkg/config"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		enforcer, err := casbin.NewEnforcer(config.C.Casbin.Model, config.C.Casbin.Policy)
		if err != nil {
			log.Fatal("enforcer not initialized, ", err)
			return
		}

		claims, err := extractClaims(accessToken, []byte(config.C.JWT.SigningKey))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		role := claims["role"]
		fmt.Println(role, c.Request.URL.String(), c.Request.Method)
		ok, err := enforcer.Enforce(role, c.Request.URL.String(), c.Request.Method)
		if err != nil {
			log.Println("could not enforce:", err)
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, map[string]string{
				"error": "user not allowed, there is a problem with authorization",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
