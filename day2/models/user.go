package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Password   string `json:"-"`
}

type UserRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Fullname string `json:"fullname,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
