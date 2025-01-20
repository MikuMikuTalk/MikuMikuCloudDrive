package services

import (
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func GetClaimsFromContext(c *gin.Context) *jwts.CustomClaims {
	claims := c.MustGet("claims").(*jwts.CustomClaims)
	return claims
}

func GetServiceContextFromContext(c *gin.Context) *ServiceContext {
	svc := c.MustGet("svc").(*ServiceContext)
	return svc
}
