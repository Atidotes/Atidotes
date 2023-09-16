package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录接口
func Login(ctx *gin.Context) {
	// a, _ := ctx.GetRawData()

	// var m map[string]interface{}
	// // 反序列化
	// _ = json.Unmarshal(a, &m)

	// fmt.Printf("userData: %v\n", string(a))
	// fmt.Printf("m: %v\n", m["user"])

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
