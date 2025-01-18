package directory_service

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/directory_types"
	"MikuMikuCloudDrive/utils/jwts"
)

// TODO: 删除目录
func (s *DirectoryService) DeleteDirectory(directoryDeleteReq directory_types.DeleteDirectoryRequest) (*directory_types.DeleteDirectoryResponse, error) {
	token := directoryDeleteReq.Token
	authConfiguration := config.ReadAuthConfig()
	claims, err := jwts.ParseJwtToken(token, authConfiguration.AuthSecret)
	if err != nil {
		return nil, err
	}
	userID := claims.UserID
	// 删除目录逻辑 删除目录前要删除目录下的所有文件
	err = s.DB.Where("user_id = ? AND directory_id = ?", userID, directoryDeleteReq.DirectoryID).Delete(&models.FileModel{}).Error
	if err != nil {
		return nil, err
	}
	err = s.DB.Where("user_id = ? AND id = ?", userID, directoryDeleteReq.DirectoryID).Delete(&models.DirectoryModel{}).Error
	if err != nil {
		return nil, err
	}
	return &directory_types.DeleteDirectoryResponse{}, nil
}
