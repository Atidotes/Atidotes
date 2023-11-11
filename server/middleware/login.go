package middleware

import (
	"fmt"
	"server/config"
	MySql "server/dao"
	"server/setting"
	"server/utils"
	"server/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取邮箱验证码
func EmailCaptcha(ctx *gin.Context) (mailbox string) {
	mailbox = ctx.Query("mailbox")
	return mailbox
}

// 登录用户
func MailboxLogin(user *setting.User) (flag bool, message string) {
	// 获取配置文件
	v, _ := config.LoadConfigFromYaml()
	tokenSecretKey := &setting.TokenSecretKey{}
	err := v.UnmarshalKey("tokenSecretKey", tokenSecretKey)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 获取路由参数
	result := MySql.DB.Debug().First(&user, "mailbox = ?", user.Mailbox)

	fmt.Printf("result: %v\n", user)

	// 验证验证码是否正确
	if result.RowsAffected == 0 {
		captchaFlag := utils.VerifyCaptcha(user.CaptchaID, user.Captcha)
		chartCaptchaFlag := utils.VerifyCaptcha(user.ChartCaptchaID, user.ChartCaptcha)

		// 插入数据
		if captchaFlag && chartCaptchaFlag {
			user.Account = utils.GenerateAccount()

			res := MySql.DB.Debug().Omit("Mobile", "Password").Create(&user)
			if res.Error != nil {
				flag = false
				message = "创建数据失败"
			} else {
				// 生成token
				token, err := jwt.GenerateToken(int(user.ID), user.Account, tokenSecretKey.SecretKey)
				if err != nil {
					flag = false
					message = "token生成失败"
				} else {
					flag = true
					message = token
				}
			}
		}
	} else {
		// 生成token
		token, err := jwt.GenerateToken(int(user.ID), user.Account, tokenSecretKey.SecretKey)
		if err != nil {
			flag = false
			message = "token生成失败"
		} else {
			flag = true
			message = token
		}
	}
	return flag, message
}

// 获取用户信息
func UserInfo(ctx *gin.Context) *jwt.Claims {
	// 获取token密钥
	v, _ := config.LoadConfigFromYaml()
	tokenSecretKey := &setting.TokenSecretKey{}
	err := v.UnmarshalKey("tokenSecretKey", tokenSecretKey)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 获取token信息
	authString := ctx.Request.Header.Get("Authorization")

	// 切割token信息
	kv := strings.Split(authString, " ")

	// 解密token
	token, _ := jwt.DecryptToken(kv[1], tokenSecretKey.SecretKey)

	fmt.Printf("token: %v\n", token)

	return token
}
