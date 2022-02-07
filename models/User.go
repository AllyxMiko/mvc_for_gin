package models

import "time"

// 为了便于演示，密码采用明文保存

type User struct {
	ID        int       `json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"create"`
}
