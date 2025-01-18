package go_test

import (
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/services/directory_service"
	"MikuMikuCloudDrive/types/directory_types"
	"testing"
)

func TestGetDirectoryInfo(t *testing.T) {
	directoryID := uint(1)
	db := core.InitGorm()
	service := directory_service.NewDirectoryService(db)
	var req directory_types.GetDirectoryInfoRequest = directory_types.GetDirectoryInfoRequest{
		DirectoryID: directoryID,
		UserID:      2,
	}
	respons, err := service.GetDirectoryInfo(req)
	if err != nil {
		t.Errorf("TestGetDirectoryInfo failed: %v", err)
	}
	t.Logf("TestGetDirectoryInfo success: %v", respons)

}
