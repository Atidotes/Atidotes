package main

import (
	"fmt"
	"io"
	"os"
	"server/config"
	MySql "server/dao"
	Models "server/models"
	routers "server/routes"
	"server/setting"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	MySQLConfig := &setting.MySQLConfig{}

	// 配置文件
	v, _ := config.LoadConfigFromYaml()
	address := v.GetString("server.address")

	// 读取配置
	err := v.UnmarshalKey("mysql", MySQLConfig)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 连接数据库
	err = MySql.InitMySQL(MySQLConfig)
	if err != nil {
		fmt.Println("数据库连接失败!!!")
		panic(err.Error())
	}
	defer MySql.CloseMySQL(MySql.DB)

	// 模型绑定
	Models.Models()

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	utils.SendEmailCode("2392503436@qq.com", "123456")

	// 路由
	r := routers.SetupRouter()
	if err := r.Run(address); err != nil {
		fmt.Println("startup service failed:%w", err)
	}
}
