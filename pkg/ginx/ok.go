package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Success",
		"data": data,
	})
}

func Fail(ctx *gin.Context, msg error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg.Error(),
		"data": nil,
	})
}

func AuthFailed(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code": 1,
		"msg":  "Authentication Failed",
		"data": nil,
	})
}

func NoAccess(ctx *gin.Context) {
	ctx.JSON(http.StatusMovedPermanently, gin.H{
		"code": 1,
		"msg":  "No access",
		"data": nil,
	})
}

func BeyondRateLimit(ctx *gin.Context) {
	ctx.JSON(http.StatusTooManyRequests, gin.H{
		"code": 1,
		"msg":  "超出访问频率限制",
		"data": nil,
	})
}
