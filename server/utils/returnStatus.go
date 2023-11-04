package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, message string, result interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"message":   message,
		"result":    result,
		"success":   true,
		"timeStamp": time.Now().Unix(),
	})
}

func Error(ctx *gin.Context, message string, result interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":      500,
		"message":   message,
		"result":    result,
		"success":   false,
		"timeStamp": time.Now().Unix(),
	})
}
