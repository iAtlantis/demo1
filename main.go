package main

import (
	"gin/demo1/biz/dal"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	dal.Init()
	SetupRouter(gin)
	gin.Run(":8080")
}
