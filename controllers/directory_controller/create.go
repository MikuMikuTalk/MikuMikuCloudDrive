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

// @Summary 创建目录 API(创建目录)
// @Description 创建目录接口
// @Tags 目录管理
// @Accept json
// @Produce json
// @Param body body directory_types.CreateDirectoryRequest true "创建目录请求参数"
// @Success 200 {object} response.Response{data=directory_types.CreateDirectoryResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /directory/create [post]
func CreateDirectory(c *gin.Context) {
	resp := response.NewResponse()
	svc := c.MustGet("svc").(*services.ServiceContext)
	dirService := directory_service.NewDirectoryService(svc.DB)

	var req directory_types.CreateDirectoryRequest
	err := c.ShouldBindHeader(&req)
	if err != nil {
		logrus.Error("绑定header失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定header失败")
		return
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Error("绑定json失败:", err)
		resp.ErrorResponse(c, http.StatusBadRequest, "绑定json失败")
		return
	}
	createDirReq := directory_types.CreateDirectoryRequest{
		Token:    req.Token,
		Name:     req.Name,
		ParentID: req.ParentID,
	}
	createResp, err := dirService.CreateDirectory(createDirReq)
	if err != nil {
		logrus.Error("创建目录失败:", err)
		resp.ErrorResponse(c, http.StatusBadGateway, "创建目录失败")
		return
	}
	resp.SuccessResponse(c, createResp, "success")
}
