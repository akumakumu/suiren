package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"-"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
