package setting

import "mvc_for_gin/models"

// 在此处写入需要迁移的表
var TableModels = []interface{}{
	&models.UserCDF{},
}
