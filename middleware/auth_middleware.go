package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := response.NewResponse()
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			resp.ErrorResponse(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ") // 去掉前缀
		authConfiguration := config.ReadAuthConfig()
		claims, err := jwts.ParseJwtToken(tokenString, authConfiguration.AuthSecret)
		if err != nil {
			resp.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		jti := claims.RegisteredClaims.ID
		username := claims.UserName
		blackList := fmt.Sprintf("blacklist_%s_%s", jti, username)
		svc := c.MustGet("svc").(*services.ServiceContext)
		rdb := svc.RedisClient
		val, err := rdb.Get(context.Background(), blackList).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			resp.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error: Failed to check blacklist")
			c.Abort()
			return
		}
		if val == "1" {
			resp.ErrorResponse(c, http.StatusUnauthorized, "登录信息已失效")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		logrus.Info("用户认证成功")
		c.Next()
	}
}
