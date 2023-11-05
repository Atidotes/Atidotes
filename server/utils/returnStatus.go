package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 200 正常请求
func Success(ctx *gin.Context, message string, result interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"message":   message,
		"result":    result,
		"success":   true,
		"timeStamp": time.Now().Unix(),
	})
}

// 404 找不到资源
func NoFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code":      404,
		"message":   "No Found",
		"result":    nil,
		"success":   false,
		"timeStamp": time.Now().Unix(),
	})
}

// 500 服务器错误
func Error(ctx *gin.Context, message string, result interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":      500,
		"message":   message,
		"result":    result,
		"success":   false,
		"timeStamp": time.Now().Unix(),
	})
}
