package file_controller

import (
	"net/http"
	"strconv"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/file_service"
	"MikuMikuCloudDrive/types/chunk_process_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Upload(c *gin.Context) {
	resp := response.NewResponse()
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Info("从form中提取file失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "error")
		return
	}
	chunkIndex, _ := strconv.Atoi(c.PostForm("chunkIndex"))
	totalChunks, _ := strconv.Atoi(c.PostForm("totalChunks"))
	fileMd5 := c.PostForm("fileMD5")
	var req chunk_process_types.ChunkUploadRequest = chunk_process_types.ChunkUploadRequest{
		File:        file,
		ChunkIndex:  chunkIndex,
		TotalChunks: totalChunks,
		FileMD5:     fileMd5,
	}
	svc := services.GetServiceContextFromContext(c)
	fileService := file_service.NewFileService(svc.DB, svc.RedisClient)
	uploadedChunksResp, err := fileService.Upload(req)
	if err != nil {
		logrus.Error("上传切片失败: ", err)
		resp.ErrorResponse(c, http.StatusInternalServerError, "error")
		return
	}
	resp.SuccessResponse(c, uploadedChunksResp, "success")
}
