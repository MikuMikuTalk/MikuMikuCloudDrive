package middleware

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/utils/jwts"
	"net/http"
	"strings"

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
		c.Set("claims", claims)
		logrus.Info("用户认证成功")
		c.Next()
	}
}
