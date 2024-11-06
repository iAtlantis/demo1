package main

import (
	"gin/demo1/biz/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g.POST("/user/:name", user.UserRegister)

}
