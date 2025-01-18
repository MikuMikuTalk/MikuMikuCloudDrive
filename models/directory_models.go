package models

import (
	"gorm.io/gorm"
)

type DirectoryModel struct {
	gorm.Model
	UserID   uint   `gorm:"index;not null"`                     // 所属用户ID
	Name     string `gorm:"size:255;uniqueIndex:idx_user_name"` // 目录名称 (unique per user)
	ParentID *uint  `gorm:"index;default:null"`                 // 父目录ID，可以为空（表示根目录）
	Path     string `gorm:"size:1024;not null"`                 // 目录路径
	// 关系
	User     UserModel        `gorm:"ForeignKey:UserID"`      // 属于哪个用户
	Parent   *DirectoryModel  `gorm:"ForeignKey:ParentID"`    // 父目录
	Children []DirectoryModel `gorm:"ForeignKey:ParentID"`    // 子目录
	Files    []FileModel      `gorm:"ForeignKey:DirectoryID"` // 目录下的文件
}

// GetUserRootDirectory 获取用户根目录
func GetUserRootDirctory(db *gorm.DB, userID uint) (*DirectoryModel, error) {
	var dir DirectoryModel
	err := db.Where("user_id = ? AND parent_id IS NULL", userID).First(&dir).Error
	return &dir, err
}

// GetSubDirectory 获取子目录列表
func GetSubDirectory(db *gorm.DB, parentID uint) ([]DirectoryModel, error) {
	var dirs []DirectoryModel
	err := db.Where("parent_id = ?", parentID).Find(&dirs).Error
	return dirs, err
}
