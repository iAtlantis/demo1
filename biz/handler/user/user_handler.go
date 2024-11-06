package user

import (
	errno "gin/demo1/pkg/error"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @router /user/:name [post]
func UserRegister(ctx *gin.Context) {
	ctx.JSON(errno.SuccessCode, gin.H{"message": "user register"})
}
