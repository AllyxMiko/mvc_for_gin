package libs

import (
	"mvc_for_gin/models"

	"gorm.io/gorm"
)

// 在这里进行表迁移
func DataBaseAutoMirgrate(d *gorm.DB) {
	// 进行表迁移
	d.AutoMigrate(&models.UserCD{})
}
