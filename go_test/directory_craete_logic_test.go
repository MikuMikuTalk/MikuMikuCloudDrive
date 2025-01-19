package go_test

import (
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/services/directory_service"
	"MikuMikuCloudDrive/types/directory_types"
	"fmt"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	db := core.InitGorm()
	// rdb := core.InitRedis()

	directory_svc := directory_service.NewDirectoryService(db)
	path := "/homes"
	req := directory_types.CreateDirectoryRequest{
		Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJuYW1lIjoidmVkYWw5ODciLCJlbWFpbCI6InZlZGFsOTg3QGV4YW1wbGUuY29tIiwiZXhwIjoxNzM3MzM1NTQwLCJqdGkiOiI4MjM1NWZjMS1hMmU4LTQ4MzktOGNlMi1iNTVhZDZhNjJhOTEifQ.SSzBNRY8zEagL_O162f2215K7RTj9rOFlUwa1Rrr6TI",
		Name:  "sss",
		Path:  &path,
	}
	res, err := directory_svc.CreateDirectory(req)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(res)
}
