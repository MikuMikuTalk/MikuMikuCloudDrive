package directory_service

import (
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/directory_types"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/sirupsen/logrus"
)

func (s *DirectoryService) CreateDirectory(directoryCreateReq directory_types.CreateDirectoryRequest, claims *jwts.CustomClaims) (*directory_types.CreateDirectoryResponse, error) {

	var parentDir models.DirectoryModel
	// 如果为空，就创建根目录
	// 如果父目录ID不为空，则获取父目录
	if directoryCreateReq.ParentID != nil {
		err := s.DB.Take(&parentDir, directoryCreateReq.ParentID).Error
		if err != nil {
			logrus.Error("获取父目录失败:", err)
			return nil, err
		}
	}

	dir := &models.DirectoryModel{
		UserID:   claims.UserID,
		Name:     directoryCreateReq.Name,
		ParentID: directoryCreateReq.ParentID,
		Path:     parentDir.Path + "/" + directoryCreateReq.Name,
	}

	err := s.DB.Create(dir).Error
	return nil, err
}
