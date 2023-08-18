package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadConfigFromYaml() (v *viper.Viper, err error) {
	dir, _ := os.Getwd() // 读取文件路径
	mode := gin.Mode()   // 获取当前环境
	fmt.Printf("mode: %v\n", mode)

	// 创建viper
	v = viper.New()

	// 根据环境读取不同配置文件
	if mode == gin.DebugMode {
		fmt.Println("当前为开发环境")
		v.SetConfigFile(dir + "/config/config-dev.yaml")
	} else if mode == gin.ReleaseMode {
		fmt.Println("当前为生产环境")
		v.SetConfigFile(dir + "/config/config-pro.yaml")
	}
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 监控配置和重新获取配置
	v.WatchConfig()

	// 监控变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return v, err
}
