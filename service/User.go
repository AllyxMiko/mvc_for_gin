package service

import (
	"fmt"
	"mvc_for_gin/models"

	"gorm.io/gorm"
)

func CheckUser(db *gorm.DB, username string) bool {
	if err := db.Select("username").Where("username=?", username).Take(&models.User{}).Error; err != nil {
		fmt.Println("check_user", err)
		return false
	}
	return true
}

func AddUser(db *gorm.DB, user *models.User) bool {
	if err := db.Create(user).Error; err != nil {
		fmt.Println("add_user", err)
		return false
	}
	fmt.Println(user.ID)
	return true
}

func GetUser(db *gorm.DB, username string) (user *models.User) {
	if err := db.Where("username=?", username).Take(&user).Error; err != nil {
		fmt.Println("get_user", err)
		return nil
	}
	return
}
