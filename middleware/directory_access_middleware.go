package middleware

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CheckDirectoryAccess(c *gin.Context) {
	logrus.Info("检查目录访问中间件调用：")
	// 获取请求路径
	requestPath := c.Request.URL.Path
	// 拼接本地文件路径
	localPath := filepath.Join("./uploads", requestPath[len("/uploads/"):])
	info, err := os.Stat(localPath)
	if err != nil {
		logrus.Error("文件不存在:", err)
		c.AbortWithStatus(http.StatusNotFound)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"Path": requestPath,
		})
		return
	}
	if info.IsDir() {
		logrus.Error("文件是目录")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}
