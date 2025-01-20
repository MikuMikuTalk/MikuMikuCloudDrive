package go_test

// func TestGetDirectoryInfo(t *testing.T) {
// 	directoryID := uint(9)
// 	db := core.InitGorm()
// 	service := directory_service.NewDirectoryService(db)
// 	var req directory_types.GetDirectoryInfoRequest = directory_types.GetDirectoryInfoRequest{
// 		DirectoryID: directoryID,
// 		Token:       "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJuYW1lIjoidmVkYWw5ODciLCJlbWFpbCI6InZlZGFsOTg3QGV4YW1wbGUuY29tIiwiZXhwIjoxNzM3MzMxODA1LCJqdGkiOiJhYzdjOWFkOC1hN2ZhLTQyNWQtYmRjMi1iM2EyODRkMzY3YTUifQ.EUECBFAR2gjOPx3jakw1SCfZjLzA5kjK2HlNHcqyCSc",
// 	}

// 	// 添加分隔线
// 	fmt.Println("\n==================== 开始测试 GetDirectoryInfo ====================")

// 	respons, err := service.GetDirectoryInfo(req)
// 	if err != nil {
// 		t.Errorf("❌ 测试失败: %v", err)
// 		fmt.Println("================================================================")
// 		return
// 	}

// 	// 美化输出格式
// 	fmt.Println("✅ 测试成功")

// 	fmt.Println("📁 目录信息:")
// 	fmt.Println(structlog.LogStruct(respons.DirectoryInfo))
// 	// fmt.Printf("  - ID: %d\n", respons.DirectoryInfo.ID)
// 	// fmt.Printf("  - 名称: %s\n", respons.DirectoryInfo.Name)
// 	// fmt.Printf("  - 路径: %s\n", respons.DirectoryInfo.Path)
// 	// fmt.Printf("  - 文件总数: %d\n", respons.DirectoryInfo.TotalFiles)
// 	// fmt.Printf("  - 总大小: %.2f MB\n", float64(respons.DirectoryInfo.TotalSize)/1024/1024)
// 	// fmt.Printf("  - 创建时间: %s\n", respons.DirectoryInfo.CreatedAt.Format("2006-01-02 15:04:05"))
// 	// fmt.Printf("  - 更新时间: %s\n", respons.DirectoryInfo.UpdatedAt.Format("2006-01-02 15:04:05"))

// 	fmt.Println("\n📂 目录内容:")

// 	for _, item := range respons.Contents {
// 		fmt.Println(structlog.LogStruct(item))
// 		// fmt.Printf("  %d. %s (%s)\n", i+1, item.Name, item.Type)
// 		// fmt.Printf("     - 大小: %.2f MB\n", float64(item.Size)/1024/1024)
// 		// fmt.Printf("     - 创建时间: %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
// 		// fmt.Printf("     - 更新时间: %s\n", item.UpdatedAt.Format("2006-01-02 15:04:05"))
// 	}

// 	fmt.Println("================================================================")
// }
