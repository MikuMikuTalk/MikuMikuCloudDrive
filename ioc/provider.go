package ioc

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/services"

	"github.com/google/wire"
)

// CoreSet 提供数据库和 Redis 实例
var CoreSet = wire.NewSet(
	core.InitGorm,  // 提供 *gorm.DB
	core.InitRedis, // 提供 *redis.Client
)

// ServiceSet 提供 ServiceContext 实例
var ServiceSet = wire.NewSet(
	services.NewServiceContext, // 提供 *services.ServiceContext
)

// ConfigSet提供ConfigContext实例
var ConfigSet = wire.NewSet(
	config.ReadAllConfig,
)
