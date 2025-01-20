package utils

import (
	"MikuMikuCloudDrive/models"
	"io"
	"os"
)

func CalculateTotalSize(files []models.FileModel) int64 {
	var total int64
	for _, file := range files {
		total += file.FileSize
	}
	return total
}
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}
