package file_controller

import "github.com/gin-gonic/gin"

func Download(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
