//go:build wireinject

package ioc

import (
	"MikuMikuCloudDrive/services"

	"github.com/google/wire"
)

// InitializeApp 生成 ServiceContext 和 AppConfig 实例
func InitializeApp() (*services.ServiceContext, error) {
	wire.Build(
		CoreSet,    // 提供 *gorm.DB 和 *redis.Client
		ConfigSet,  // 提供 *config.AppConfig
		ServiceSet, // 提供 *services.ServiceContext
	)
	return &services.ServiceContext{}, nil
}
