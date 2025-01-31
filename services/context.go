package services

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	DB            *gorm.DB
	RedisClient   *redis.Client
	Configuration *config.Config
}

func NewServiceContext(db *gorm.DB, rdb *redis.Client, config *config.Config) *ServiceContext {
	return &ServiceContext{
		DB:            db,
		RedisClient:   rdb,
		Configuration: config,
	}
}
func GetClaimsFromContext(c *gin.Context) *jwts.CustomClaims {
	claims := c.MustGet("claims").(*jwts.CustomClaims)
	return claims
}

func GetServiceContextFromContext(c *gin.Context) *ServiceContext {
	svc := c.MustGet("svc").(*ServiceContext)
	return svc
}
