package routers

import (
	"server/controllers"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件
	r.Static("/static", "static")
	// 404处理
	r.NoRoute(func(ctx *gin.Context) {
		utils.NoFound(ctx)
	})

	// v1
	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/system/chartCaptcha", controllers.ChartCaptcha)
		v1Group.GET("/system/emailCaptcha", controllers.EmailCaptcha)
		v1Group.POST("/system/login", controllers.Login)
	}
	return r
}
