package setting

import (
	"io/ioutil"

	"github.com/spf13/viper"
)

type DataBaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Passwd   string `mapstructure:"passwd"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type ProjectConfig struct {
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type Configs struct {
	ProjectConfig  ProjectConfig  `mapstructure:"project"`
	DataBaseConfig DataBaseConfig `mapstructure:"database"`
}

type ConfigManager struct {
	Config  *viper.Viper
	Configs Configs
}

var Cfm = new(ConfigManager)
var DbConfigs = Cfm.Configs.DataBaseConfig
var PjtConfigs = Cfm.Configs.ProjectConfig

func init() {
	Cfm.Config = viper.New()
	Cfm.Config.SetConfigFile("./configs/config.yaml")
	Cfm.Config.ReadInConfig()
	Cfm.Config.Unmarshal(&Cfm.Configs)
}

func CreateConfigsFile() {
	demoFile := []byte(`# 项目的配置项
project:
  # web服务的地址
  host: "localhost"
  # web服务的端口号
  port: 8080
  # 是否开启debug模式
  debug: true

# 数据库配置项，目前仅支持mysql数据库
database:
  # 数据库主机
  host: "127.0.0.1"
  # 数据库用户名
  user: "root"
  # 数据库密码
  passwd: "123456"
  # 数据库端口
  port: 3306
  # 数据库名称
  database: "demo"
  # 数据库编码
  charset: "utf8"`)
	ioutil.WriteFile("./configs/config.yaml", demoFile, 0777)
}
