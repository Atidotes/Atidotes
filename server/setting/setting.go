package setting

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Root     string `yaml:"root"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
}

// 邮箱配置
type Mailbox struct {
	Sender   string `yaml:"sender"`
	AuthCode string `yaml:"authCode"`
}
