package service

import (
	db "mvc_for_gin/database"
	"mvc_for_gin/models"
)

func AddUser(u *models.User) (err error) {
	if err = db.DbConn.Create(u).Error; err != nil {
		return err
	}
	return
}

func FindAll() (Users []*models.User, err error) {
	if err = db.DbConn.Find(&Users).Error; err != nil {
		return nil, err
	}
	return
}

func DelUser(id int) int64 {
	return db.DbConn.Where("id=?", id).Delete(&models.User{}).RowsAffected
}
