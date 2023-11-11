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

	// 发送验证码
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
			flag, message := middleware.MailboxLogin(&user)
			if flag {
				// 设置响应头
				ctx.Writer.Header().Set("Authorization", message)
				// 返回数据
				utils.Success(ctx, "邮箱验证码登录成功", message)
			} else {
				utils.Error(ctx, message, message)
			}
		} else if user.LoginType == 2 {
			utils.Success(ctx, "账号密码登录成功", nil)
		}
	}
}

// 获取用户信息
func UserInfo(ctx *gin.Context) {
	token := middleware.UserInfo(ctx)
	utils.Success(ctx, "获取成功", token)
}
