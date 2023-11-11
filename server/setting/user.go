package setting

import (
	"gorm.io/gorm"
)

// 用户接口
type User struct {
	gorm.Model
	LoginType      int    `gorm:"-"`                 // 账号类型 0 手机验证码登录 1 邮箱验证码登录 2 账号密码登录
	Mobile         string `gorm:"type:int;"`         // 手机号
	Mailbox        string `gorm:"type:varchar(32);"` // 邮箱号
	Account        string `gorm:"type:string;"`      // 用户账号
	Password       string `gorm:"type:varchar(32)"`  // 用户密码
	Captcha        string `gorm:"-"`                 // 验证码
	CaptchaID      string `gorm:"-"`                 // 验证码ID
	ChartCaptcha   string `gorm:"-"`                 // 图形验证码
	ChartCaptchaID string `gorm:"-"`                 // 图形验证码ID
}
