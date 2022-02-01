package libs

import (
	"log"
	"mvc_for_gin/configs"
	"mvc_for_gin/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB
var dbError error

// 初始化数据库
func InitDataBase() {
	DbConn, dbError = gorm.Open(mysql.Open(connectUrl()), &gorm.Config{})
	if dbError != nil {
		log.Fatal("数据库配置出现错误！请检查配置！如果您不需要使用到数据库请在main函数中配置.NoDataBase()以此跳过数据库配置。")
	}
	// 进行数据表迁移
	DataBaseAutoMirgrates(DbConn, setting.TableModels...)
}

func connectUrl() string {
	return configs.USER + ":" + configs.PASSWD + "@tcp(" + configs.HOST + ":" + configs.PORT + ")/" + configs.DATABASE + "?charset=utf8&parseTime=True&loc=Local"
}
