package user_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"MikuMikuCloudDrive/types/logout_types"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func (s *UserService) Logout(logoutReq logout_types.LogoutRequest, claims *jwts.CustomClaims) (*logout_types.LogoutResponse, error) {

	now := time.Now()
	jti := claims.RegisteredClaims.ID
	userName := claims.UserName
	expiration := claims.ExpiresAt.Time.Sub(now)
	result, err := s.RedisClient.Get(context.Background(), "blacklist_"+jti+"_"+userName).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if result == "1" {
		return nil, errors.New("您已经注销过了")
	}
	err = s.RedisClient.Set(context.Background(), "blacklist_"+jti+"_"+userName, true, expiration).Err()
	if err != nil {
		errInfo := fmt.Errorf("用户注销信息存储失败:%v", err)
		logrus.Error(errInfo)
		return nil, err
	}
	return &logout_types.LogoutResponse{}, nil
}
