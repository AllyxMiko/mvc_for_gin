package setting

import (
	"io/ioutil"
	"log"
	"os"

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
	ProjectConfigs  ProjectConfig  `mapstructure:"project"`
	DataBaseConfigs DataBaseConfig `mapstructure:"database"`
}

type ConfigManager struct {
	Config  *viper.Viper
	Configs Configs
}

var Cfm = new(ConfigManager)
var DbConfigs = &Cfm.Configs.DataBaseConfigs
var PjtConfigs = &Cfm.Configs.ProjectConfigs

func init() {
	Cfm.Config = viper.New()
	Cfm.Config.SetConfigFile("./configs/config.yaml")
	Cfm.Config.ReadInConfig()
	Cfm.Config.Unmarshal(&Cfm.Configs)
}

func CheckConfigs() {
	_, err := os.Stat("./configs/config.yaml")
	if err != nil {
		_, err = os.Stat("./configs")
		if err != nil {
			os.Mkdir("./configs", 0755)
		}
		CreateConfigsFile()
		log.Fatal("配置文件不存在，系统将自动生成！请按照说明修改配置文件！然后重新启动项目！")
	}
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
	ioutil.WriteFile("./configs/config.yaml", demoFile, 0755)
}
