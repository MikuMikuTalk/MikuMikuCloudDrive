package file_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/file_service"
	"MikuMikuCloudDrive/types/file_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Download(c *gin.Context) {
	logrus.Debug("下载文件服务调用")
	resp := response.NewResponse()

	downloadFileRequest := file_types.DownloadFileRequest{
		FileID: c.Param("file_id"),
	}
	svc := services.GetServiceContextFromContext(c)
	db := svc.DB
	rdb := svc.RedisClient
	fileService := file_service.NewFileService(db, rdb)
	// 获取claims
	claims := services.GetClaimsFromContext(c)
	// 获取文件服务
	downloadResponse, err := fileService.DownloadFile(downloadFileRequest, claims)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	resp.SuccessResponse(c, downloadResponse, "success")
}
