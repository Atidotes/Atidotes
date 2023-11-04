package routers

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件
	r.Static("/static", "static")

	// v1
	v1Group := r.Group("v1")
	{
		v1Group.GET("/system/emailCaptcha", controllers.EmailCaptcha)
		v1Group.POST("/system/login", controllers.Login)
	}
	return r
}
