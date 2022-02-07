package database

import (
	"log"
	"mvc_for_gin/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB
var dbError error

// 初始化数据库
func InitDataBase() {
	DbConn, dbError = gorm.Open(mysql.Open(connectUrl()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if dbError != nil {
		log.Fatal("数据库配置出现错误！请检查配置！如果您不需要使用到数据库请在main函数中配置.NoDataBase()以此跳过数据库配置。")
	}
	// 进行数据表迁移
	DataBaseAutoMirgrates(DbConn, setting.TableModels...)
}

func connectUrl() string {
	return setting.DbConfigs.User +
		":" +
		setting.DbConfigs.Passwd +
		"@tcp(" +
		setting.DbConfigs.Host +
		":" +
		setting.DbConfigs.Port + ")/" +
		setting.DbConfigs.Database +
		"?charset=" + setting.DbConfigs.Charset +
		"&parseTime=True&loc=Local"
}

func DataBaseAutoMirgrates(d *gorm.DB, dst ...interface{}) {
	// 进行表迁移
	d.AutoMigrate(dst...)
}
