package file_models

import "gorm.io/gorm"

type FileModel struct {
	gorm.Model
	UserID   uint
	FileName string
	FilePath string
	FileSize int64
	FileHash string
}
