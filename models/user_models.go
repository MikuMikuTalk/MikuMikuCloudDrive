package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName string `gorm:"unique;not null"` // 用户名，唯一
	Password string `gorm:"not null"`        // 密码
	Avatar   string // 头像路径
	Email    string `gorm:"unique;not null"` // 邮箱，唯一
	// 一个用户有多个目录和文件
	Directories []DirectoryModel `gorm:"ForeignKey:UserID"`
	Files       []FileModel      `gorm:"ForeignKey:UserID"`
}
