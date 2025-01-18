package directory_controller

import "github.com/gin-gonic/gin"

// @Summary 获取目录内容 API(获取目录内容)
// @Description 获取目录内容接口
// @Tags 目录管理
// @Accept json
// @Produce json
// @Param body body directory_types.CreateDirectoryRequest true "创建目录请求参数"
// @Success 200 {object} response.Response{data=directory_types.CreateDirectoryResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /directory/content [get]
func GetDirectoryContent(c *gin.Context) {
	// TODO: 实现获取目录内容逻辑
}
