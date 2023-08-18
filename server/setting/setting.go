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

type User struct {
	gorm.Model
	User     string `gorm:"type:varchar(100);not null;unique_index;"`
	Password string `gorm:"type:varchar(100);not null"`
	Mailbox  string `gorm:"type:varchar(100);not null"`
}
