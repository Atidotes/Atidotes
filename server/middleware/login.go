package middleware

import (
	"github.com/gin-gonic/gin"
)

// 获取邮箱验证码
func EmailCaptcha(ctx *gin.Context) (mailbox string) {
	// var user setting.User
	mailbox = ctx.Query("mailbox")
	// result := MySql.DB.Debug().First(&user, "mailbox = ?", mailbox)
	return mailbox
}
