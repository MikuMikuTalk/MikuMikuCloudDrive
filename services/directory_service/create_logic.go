package directory_service

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/directory_types"
	"MikuMikuCloudDrive/utils/jwts"
)

func (s *DirectoryService) CreateDirectory(directoryCreateReq directory_types.CreateDirectoryRequest) (*directory_types.CreateDirectoryResponse, error) {
	authConfiguration := config.ReadAuthConfig()
	claims, err := jwts.ParseJwtToken(directoryCreateReq.Token, authConfiguration.AuthSecret)
	if err != nil {
		return nil, err
	}
	dir := &models.DirectoryModel{
		UserID:   claims.UserID,
		Name:     directoryCreateReq.Name,
		ParentID: &directoryCreateReq.ParentID,
	}
	err = s.DB.Create(dir).Error
	return nil, err
}
