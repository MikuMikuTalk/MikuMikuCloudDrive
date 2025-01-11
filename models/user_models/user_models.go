package user_models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
	Avatar   string
	Email    string `gorm:"unique"`
}
