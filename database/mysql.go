package libs

import (
	"mvc_for_gin/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB
var dbError error

// 初始化数据库
func InitDataBase() {
	DbConn, dbError = gorm.Open(mysql.Open(connectUrl()), &gorm.Config{})
	if dbError != nil {
		panic(dbError)
	}

	DataBaseAutoMirgrate(DbConn)
}

func connectUrl() string {
	return configs.USER + ":" + configs.PASSWD + "@tcp(" + configs.HOST + ":" + configs.PORT + ")/" + configs.DATABASE + "?charset=utf8&parseTime=True&loc=Local"
}
