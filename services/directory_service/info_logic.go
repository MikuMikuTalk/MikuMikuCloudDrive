package directory_service

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/directory_types"
	"MikuMikuCloudDrive/utils/jwts"
	"fmt"

	"gorm.io/gorm"
)

type DirectoryInfoService struct {
	DB *gorm.DB
}

func (s *DirectoryService) GetDirectoryInfo(req directory_types.GetDirectoryInfoRequest) (*directory_types.GetDirectoryInfoResponse, error) {
	var directory models.DirectoryModel
	var files []models.FileModel
	var subDirs []models.DirectoryModel
	authConfiguration := config.ReadAuthConfig()
	claims, err := jwts.ParseJwtToken(req.Token, authConfiguration.AuthSecret)
	if err != nil {
		return nil, err
	}
	directoryID := req.DirectoryID

	userID := claims.UserID
	// Get directory info
	if err := s.DB.Where("id = ? and user_id = ?", directoryID, userID).First(&directory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("目录不存在")
		}
		return nil, fmt.Errorf("获取目录信息失败: %v", err)
	}

	// Get directory with preloaded files and children directories
	if err := s.DB.
		Preload("Files").
		Preload("Children").
		Where("id = ? and user_id = ?", directoryID, userID).
		First(&directory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("目录不存在")
		}
		return nil, fmt.Errorf("获取目录信息失败: %v", err)
	}

	files = directory.Files
	subDirs = directory.Children

	// Check if root directory
	isRoot := directory.ParentID == nil

	// Build response
	response := &directory_types.GetDirectoryInfoResponse{
		DirectoryInfo: directory_types.DirectoryInfo{
			ID:         directory.ID,
			Name:       directory.Name,
			Path:       directory.Path,
			CreatedAt:  directory.CreatedAt,
			UpdatedAt:  directory.UpdatedAt,
			TotalFiles: len(files),
			TotalSize:  calculateTotalSize(files),
			IsRoot:     isRoot,
		},
		Contents: buildDirectoryContents(files, subDirs),
	}

	return response, nil
}

func calculateTotalSize(files []models.FileModel) int64 {
	var total int64
	for _, file := range files {
		total += file.FileSize
	}
	return total
}

func buildDirectoryContents(files []models.FileModel, dirs []models.DirectoryModel) []directory_types.DirectoryItem {
	var contents []directory_types.DirectoryItem

	// Add files
	for _, file := range files {
		contents = append(contents, directory_types.DirectoryItem{
			ID:        file.ID,
			Name:      file.FileName,
			Type:      "file",
			Size:      file.FileSize,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
		})
	}

	// Add directories
	for _, dir := range dirs {
		contents = append(contents, directory_types.DirectoryItem{
			ID:        dir.ID,
			Name:      dir.Name,
			Type:      "directory",
			CreatedAt: dir.CreatedAt,
			UpdatedAt: dir.UpdatedAt,
		})
	}

	return contents
}
