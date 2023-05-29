package router

import (
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/api-gateway/internal/controller/user"
	"github.com/i-akbarshoh/api-gateway/internal/middleware"
)

func New(c *user.Controller) *gin.Engine {
	e := gin.Default()
	e.Use(middleware.Authorizer())

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	e.POST("/register", c.SignUp)
	e.POST("/login", c.Login)

	return e
}