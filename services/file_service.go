package services

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models/file_models"
	"MikuMikuCloudDrive/types/chunk_process_types"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
)

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{
		DB: db,
	}
}

func (s *FileService) Upload(req chunk_process_types.ChunkUploadRequest) (*chunk_process_types.ChunkUploadedResponse, error) {

	appConfig := config.ReadAppConfig()

	file := req.File
	chunkIndex := req.ChunkIndex
	fileMd5 := req.FileMD5
	totalChunks := req.TotalChunks

	cacheDir := appConfig.CacheDir

	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return nil, err
	}
	chunkFilename := fmt.Sprintf("%s_%s_%d", file.Filename, fileMd5, chunkIndex)
	cachePath := filepath.Join(cacheDir, chunkFilename)

	// 保存分片文件
	dst, err := os.Create(cachePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	logrus.Info("Received chunk: ", chunkIndex+1, " of ", totalChunks, " for file ", fileMd5)

	return &chunk_process_types.ChunkUploadedResponse{}, nil
}
func (s *FileService) GetUploadedChunks(req chunk_process_types.GetUploadedChunksRequest) (*chunk_process_types.GetUploadedChunksResponse, error) {
	fileName := req.FileName
	totalChunks := req.TotalChunks
	fileMD5 := req.FileMD5
	appConfig := config.ReadAppConfig()

	var uploadedChunks []int = []int{}
	for i := 0; i < totalChunks; i++ {
		chunkFileName := fmt.Sprintf("%s_%s_%d", fileName, fileMD5, i)
		chunkPath := filepath.Join(appConfig.CacheDir, chunkFileName)
		if _, err := os.Stat(chunkPath); err == nil {
			uploadedChunks = append(uploadedChunks, i)
		}
	}
	return &chunk_process_types.GetUploadedChunksResponse{
		ChunksArray: uploadedChunks,
	}, nil

}
func (s *FileService) MergeChunksToFile(req chunk_process_types.MergeChunksRequest) (*chunk_process_types.MergeChunksResponse, error) {
	fileName := req.FileName
	totalChunks := req.TotalChunks
	fileMD5 := req.FileMD5
	appConfig := config.ReadAppConfig()

	cacheDir := appConfig.CacheDir
	uploadedDir := appConfig.UploadDir

	// 确保上传目录存在
	if err := os.MkdirAll(uploadedDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %v", err)
	}

	// 创建最终文件
	finalFilePath := filepath.Join(uploadedDir, fileName)
	out, err := os.Create(finalFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create final file: %v", err)
	}
	defer out.Close()

	// 合并分片文件
	for i := 0; i < totalChunks; i++ {
		chunkFileName := fmt.Sprintf("%s_%s_%d", fileName, fileMD5, i)
		chunkPath := filepath.Join(cacheDir, chunkFileName)

		// 打开分片文件
		chunk, err := os.Open(chunkPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open chunk file: %v", err)
		}

		// 将分片内容写入最终文件
		if _, err := io.Copy(out, chunk); err != nil {
			chunk.Close() // 立即关闭文件句柄
			return nil, fmt.Errorf("failed to merge chunk: %v", err)
		}

		// 关闭分片文件句柄
		chunk.Close()

		// 删除分片文件
		if err := os.Remove(chunkPath); err != nil {
			return nil, fmt.Errorf("failed to delete chunk file: %v", err)
		}
		fileInfo, err := os.Stat(finalFilePath)
		if err != nil {
			logrus.Errorf("failed to stat file: %v", err)
			return nil, fmt.Errorf("无法获取文件信息:%v", err)
		}

		// 获取文件大小
		fileSize := fileInfo.Size()
		fileModel := file_models.FileModel{
			UserID:   1,
			FileName: fileName,
			FileHash: fileMD5,
			FilePath: finalFilePath,
			FileSize: fileSize,
		}
		// 创建数据库记录
		err = s.DB.Create(&fileModel).Error
		if err != nil {
			logrus.Errorf("failed to save file info to database: %v", err)
			return nil, err
		}
	}

	return &chunk_process_types.MergeChunksResponse{}, nil
}
