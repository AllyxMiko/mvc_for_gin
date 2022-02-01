package libs

import (
	"gorm.io/gorm"
)

// 在这里进行表迁移
func DataBaseAutoMirgrates(d *gorm.DB, dst ...interface{}) {
	// 进行表迁移
	d.AutoMigrate(dst...)
}
