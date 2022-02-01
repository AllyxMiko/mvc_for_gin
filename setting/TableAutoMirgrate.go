package setting

import "mvc_for_gin/models"

var TableModels = []interface{}{
	// 在此处写入需要迁移的表
	&models.UserCDF{},
}
