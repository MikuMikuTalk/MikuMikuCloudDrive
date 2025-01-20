package file_service

import (
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/types/file_types"
	"MikuMikuCloudDrive/utils/jwts"
	"MikuMikuCloudDrive/utils/url"
	"errors"

	"github.com/sirupsen/logrus"
)

func (s *FileService) DownloadFile(downloadFileRequest file_types.DownloadFileRequest, claims *jwts.CustomClaims) (*file_types.DownloadFileResponse, error) {

	fileID := downloadFileRequest.FileID
	// 获取用户ID
	userID := claims.UserID
	fileModel := models.FileModel{}

	err := s.DB.Take(&fileModel, fileID).Error // 获取文件信息
	if err != nil {
		logrus.Error("获取文件信息失败:" + err.Error())
		return nil, errors.New("获取文件信息失败:" + err.Error())
	}
	fileOwnerID := fileModel.UserID
	if fileOwnerID != userID {
		logrus.Error("用户没有权限下载这个文件")
		return nil, errors.New("用户没有权限下载这个文件")
	}
	// 获取文件路径
	filePath := fileModel.FilePath
	// 获取文件内容
	// 防止内存溢出
	fileDownloadResponse := file_types.DownloadFileResponse{
		FileID:   fileID,
		Url:      url.ConcatWebUrl(filePath),
		FileName: fileModel.FileName,
		FileHash: fileModel.FileHash,
	}

	return &fileDownloadResponse, nil
}
