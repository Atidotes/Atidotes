package setting

import (
	"gorm.io/gorm"
)

// 用户接口
type User struct {
	gorm.Model
	LoginType    int    `gorm:"type:int;not null"`               // 账号类型 0 手机验证码登录 1 邮箱验证码登录 2 账号密码登录
	Mobile       string `gorm:"type:int;not null"`               // 手机号
	Mailbox      string `gorm:"type:varchar(32);not null;"`      // 邮箱号
	Account      string `gorm:"type:int;not null;unique_index;"` // 用户账号
	Password     string `gorm:"type:varchar(32);not null"`       // 用户密码
	Captcha      string `gorm:"type:char(100);not null"`         // 验证码
	ChartCaptcha string `gorm:"type:char(100);not null"`         // 图形验证码
}
