package Models

import (
	MySql "server/dao"
	"server/setting"
)

func Models() {
	MySql.DB.AutoMigrate(&setting.User{})
}
