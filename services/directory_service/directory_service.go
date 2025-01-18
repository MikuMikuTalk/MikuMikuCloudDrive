package directory_service

import "gorm.io/gorm"

type DirectoryService struct {
	DB *gorm.DB
}

func NewDirectoryService(db *gorm.DB) *DirectoryService {
	return &DirectoryService{
		DB: db,
	}
}
