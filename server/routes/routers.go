package routers

import (
	"server/controllers"
	"server/utils"
	"server/utils/jwt"

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

	// 不需要做token验证的接口
	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/system/chartCaptcha", controllers.ChartCaptcha) // 获取图形验证码
		v1Group.GET("/system/emailCaptcha", controllers.EmailCaptcha) // 获取邮箱验证码
		v1Group.POST("/system/login", controllers.Login)              // 登录注册
	}

	// 系统接口 需做token验证
	systemGroup := r.Group("/api/v1/system", jwt.Auth())
	{
		systemGroup.GET("/userInfo", controllers.UserInfo)
	}
	return r
}
