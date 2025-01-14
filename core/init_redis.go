package core

import (
	"MikuMikuCloudDrive/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	redisConfiguration := config.ReadRedisConfig()
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfiguration.Host, redisConfiguration.Port),
		Password: redisConfiguration.Password,
		DB:       redisConfiguration.Database,
	})
	_, err := rdbClient.Ping(context.Background()).Result()
	if err != nil {
		logrus.Error("Init redis fail")
		return nil
	}
	return rdbClient
}
