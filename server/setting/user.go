package setting

import (
	"gorm.io/gorm"
)

// 用户接口
type User struct {
	gorm.Model
	User       string `gorm:"type:varchar(100);not null;unique_index;"` // 用户账号
	Password   string `gorm:"type:varchar(100);not null"`               // 用户密码
	Mailbox    string `gorm:"type:varchar(100);not null"`               // 用户邮箱
	NumberType int    `gorm:"type:int;not null"`                        // 账号类型
}
