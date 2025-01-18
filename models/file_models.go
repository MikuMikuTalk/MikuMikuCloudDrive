package models

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	UserID      uint   `gorm:"index;not null"`         // 所属用户ID
	DirectoryID uint   `gorm:"index;not null"`         // 所属目录ID
	FileName    string `gorm:"size:255;not null"`      // 文件名
	FilePath    string `gorm:"not null"`               // 文件存储路径
	FileSize    int64  `gorm:"not null"`               // 文件大小
	FileHash    string `gorm:"size:64;not null;index"` // 文件哈希值，防止重复上传
	// 关系
	User      UserModel      `gorm:"ForeignKey:UserID"`      // 属于哪个用户
	Directory DirectoryModel `gorm:"ForeignKey:DirectoryID"` // 所属目录
}

/*

默认目录：可以设置一个默认目录（如“我的文件”或“根目录”），当用户未指定目录时，文件自动存放到该目录。

自动创建目录：允许用户在上传文件时指定完整的路径，系统自动创建路径中不存在的目录，类似于文件系统中的 mkdir -p 命令。


在上传文件时，如果指定的目录不存在，系统应自动创建该目录及其上级目录（如果需要）。

提供用户友好的界面，允许用户选择或输入文件存放的目录路径。
*/
