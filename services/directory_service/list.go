package directory_service

import (
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/directory_types"
	utils "MikuMikuCloudDrive/utils/files"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/sirupsen/logrus"
)

func (s *DirectoryService) GetDirectoryList(req directory_types.GetDirectoryListRequest, claims *jwts.CustomClaims) (*directory_types.GetDirectoryListResponse, error) {
	userID := claims.UserID
	dirs := []models.DirectoryModel{}
	err := s.DB.Preload("Files").
		Preload("Children").Find(&dirs, "user_id = ?", userID).Error
	if err != nil {
		logrus.Error("查询用户目录失败:", err)
		return nil, err
	}
	var directoryInfo []directory_types.DirectoryInfo = make([]directory_types.DirectoryInfo, 0)
	for _, dir := range dirs {
		directoryInfo = append(directoryInfo, directory_types.DirectoryInfo{
			ID:          dir.ID,
			Name:        dir.Name,
			Path:        dir.Path,
			CreatedAt:   dir.CreatedAt,
			UpdatedAt:   dir.UpdatedAt,
			TotalFiles:  len(dir.Files),
			TotalSize:   utils.CalculateTotalSize(dir.Files),
			IsRoot:      dir.ParentID == nil,
			IsShared:    false,
			Permissions: "rwxr-xr-x",
		})
	}
	return &directory_types.GetDirectoryListResponse{
		Directories: directoryInfo,
	}, nil
}
