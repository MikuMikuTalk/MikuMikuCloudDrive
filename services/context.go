package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}
