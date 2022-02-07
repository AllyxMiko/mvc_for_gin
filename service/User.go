package service

import (
	"mvc_for_gin/models"

	"gorm.io/gorm"
)

func CheckUser(db *gorm.DB, username string) bool {
	if err := db.Select("username").Where("username=?", username).Take(&models.User{}).Error; err != nil {
		return false
	}
	return true
}

func AddUser(db *gorm.DB, user *models.User) bool {
	if err := db.Create(user).Error; err != nil {
		return false
	}
	return true
}

func GetUser(db *gorm.DB, username string) (user *models.User) {
	if err := db.Where("username=?", username).Take(&user).Error; err != nil {
		return nil
	}
	return
}
