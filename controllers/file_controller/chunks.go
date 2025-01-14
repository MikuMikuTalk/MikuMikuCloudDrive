package file_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/types/chunk_process_types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetUploadedChunks(c *gin.Context) {
	var getUploadedReq chunk_process_types.GetUploadedChunksRequest
	resp := response.NewResponse()
	if err := c.ShouldBindJSON(&getUploadedReq); err != nil {
		logrus.Error("绑定json失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定json失败")
		return
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	fileService := services.NewFileService(svc.DB)
	getupLoadedResponse, err := fileService.GetUploadedChunks(getUploadedReq)
	if err != nil {
		logrus.Error("获取已上传切片数失败", err)
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessResponse(c, getupLoadedResponse, "success")
}

func MergeChunks(c *gin.Context) {
	var mergeReq chunk_process_types.MergeChunksRequest
	resp := response.NewResponse()
	if err := c.ShouldBindJSON(&mergeReq); err != nil {
		logrus.Error("绑定json失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定json失败")
		return
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	fileService := services.NewFileService(svc.DB)

	mergedResp, err := fileService.MergeChunksToFile(mergeReq)
	if err != nil {
		logrus.Error("合并分片失败", err)
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessResponse(c, mergedResp, "success")
}
