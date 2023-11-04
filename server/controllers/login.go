package controllers

import (
	"net/http"
	"server/middleware"
	"server/setting"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 登录接口
func Login(ctx *gin.Context) {
	user := setting.User{}
	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
	} else {
		if user.LoginType == 0 {
			utils.Success(ctx, "手机验证码登录成功", nil)
		} else if user.LoginType == 1 {
			middleware.MailboxLogin(&user)
			utils.Success(ctx, "邮箱验证码登录成功", nil)
		} else if user.LoginType == 2 {
			utils.Success(ctx, "账号密码登录成功", nil)
		}
	}
}

func EmailCaptcha(ctx *gin.Context) {
	utils.Success(ctx, "发送成功", nil)
}
