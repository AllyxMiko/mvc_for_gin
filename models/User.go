package models

import "time"

// 为了便于演示，密码采用明文保存

type User struct {
	ID        int       `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:create" json:"create"`
}
