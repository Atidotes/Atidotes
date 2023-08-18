package MySql

import (
	"fmt"
	"server/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

// 连接MySQL
func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Root, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		}})

	fmt.Println("数据库连接成功！")
	return err
}

// 关闭MySQL
func CloseMySQL(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("DB Error!!!")
		panic(err.Error())
	}

	sqlDB.Close()
	fmt.Println("数据库连接关闭!!!")
}
