package controllers

import (
	"net/http"
	"server/middleware"
	"server/setting"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 图形验证码
func ChartCaptcha(ctx *gin.Context) {
	id, base64, _, err := utils.MakeCaptcha()
	if err != nil {
		utils.Error(ctx, "获取失败", err)
	} else {
		utils.Success(ctx, "获取成功", gin.H{
			"id":      id,
			"captcha": base64,
		})
	}
}

// 邮箱验证码
func EmailCaptcha(ctx *gin.Context) {
	mailbox := middleware.EmailCaptcha(ctx)

	id, _, code, err := utils.MakeCaptcha()
	if err != nil || mailbox == "" {
		utils.Error(ctx, "获取失败", err)
	} else {
		err2 := utils.SendEmailCode(mailbox, code)
		if err2 != nil {
			utils.Error(ctx, "获取失败", err)
		} else {
			utils.Success(ctx, "获取成功", gin.H{
				"id": id,
			})
		}
	}
}

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
			// middleware.MailboxLogin(&user)
			utils.Success(ctx, "邮箱验证码登录成功", nil)
		} else if user.LoginType == 2 {
			utils.Success(ctx, "账号密码登录成功", nil)
		}
	}
}
