package model

import "gorm.io/gorm"

type UserImage struct {
	gorm.Model
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserEmail    string `json:"user_email"`
}
