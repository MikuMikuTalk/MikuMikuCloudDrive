package user_service

import (
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type UserService struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewUserService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		DB:          db,
		RedisClient: rdb,
	}
}
