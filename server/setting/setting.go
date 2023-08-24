package setting

import "gorm.io/gorm"

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Root     string `yaml:"root"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
}

// 用户接口
type User struct {
	gorm.Model
	User     string `gorm:"type:varchar(100);not null;unique_index;"`
	Password string `gorm:"type:varchar(100);not null"`
	Mailbox  string `gorm:"type:varchar(100);not null"`
}

// 邮箱配置
type Mailbox struct {
	Sender   string `yaml:sender`
	AuthCode string `yaml:authCode`
}
