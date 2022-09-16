package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Fullname   string `json:"fullname,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"-"`
}

type UserRequest struct {
	Fullname string `json:"fullname" validate:"required,min=3,max=25"`
	Username string `json:"username" validate:"required,lowercase,min=5,max=25"`
	Password string `json:"password" validate:"required,min=10,max=255"`
}

type UserUpdateRequest struct {
	Fullname string `json:"fullname,omitempty" validate:"omitempty,min=3,max=25"`
	Username string `json:"username,omitempty" validate:"omitempty,lowercase,min=5,max=25"`
	Password string `json:"password,omitempty" validate:"omitempty,min=10,max=255"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,lowercase,min=5,max=25"`
	Password string `json:"password" validate:"required,min=10,max=255"`
}
