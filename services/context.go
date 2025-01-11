package services

import "gorm.io/gorm"

type ServiceContext struct {
	DB *gorm.DB
}
