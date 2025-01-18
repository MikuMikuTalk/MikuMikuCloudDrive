package directory_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/directory_service"
	"MikuMikuCloudDrive/types/directory_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary 删除目录 API(删除目录)
// @Description 删除目录接口
// @Tags 目录管理
// @Accept json
// @Produce json
// @Param body body directory_types.DeleteDirectoryRequest true "删除目录请求参数"
// @Success 200 {object} response.Response{data=directory_types.DeleteDirectoryResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /directory/delete [post]
func DeleteDirectory(c *gin.Context) {
	resp := response.NewResponse()
	svc := c.MustGet("svc").(*services.ServiceContext)
	dirService := directory_service.NewDirectoryService(svc.DB)
	var deleteDirectoryReq directory_types.DeleteDirectoryRequest = directory_types.DeleteDirectoryRequest{}
	err := c.ShouldBindHeader(&deleteDirectoryReq)
	if err != nil {
		logrus.Error("绑定header失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定header失败")
		return
	}
	err = c.ShouldBindJSON(&deleteDirectoryReq)
	if err != nil {
		logrus.Error("绑定json失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定json失败")
		return
	}

	dirDeleteResp, err := dirService.DeleteDirectory(deleteDirectoryReq)
	if err != nil {
		logrus.Error("删除目录失败:", err)
		resp.ErrorResponse(c, http.StatusBadGateway, "删除目录失败")
		return
	}
	resp.SuccessResponse(c, dirDeleteResp, "success")
}
