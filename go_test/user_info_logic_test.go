package go_test

// func TestUserInfo(t *testing.T) {
// 	// 添加分隔线
// 	fmt.Println("\n==================== 开始测试用户信息 ====================")

// 	db := core.InitGorm()
// 	rdb := core.InitRedis()
// 	userService := user_service.NewUserService(db, rdb)
// 	userinfoReq := userinfo_types.UserInfoRequest{
// 		Token: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJuYW1lIjoidmVkYWw5ODciLCJlbWFpbCI6InZlZGFsOTg3QGV4YW1wbGUuY29tIiwiZXhwIjoxNzM3MzMxODA1LCJqdGkiOiJhYzdjOWFkOC1hN2ZhLTQyNWQtYmRjMi1iM2EyODRkMzY3YTUifQ.EUECBFAR2gjOPx3jakw1SCfZjLzA5kjK2HlNHcqyCSc",
// 	}

// 	res, err := userService.GetUserInfo(userinfoReq)
// 	if err != nil {
// 		t.Errorf("❌ 测试失败: %v", err)
// 		fmt.Println("================================================================")
// 		return
// 	}

// 	// 使用testlogs库输出用户信息
// 	fmt.Println("✅ 测试成功")
// 	fmt.Println("👤 用户信息:")
// 	fmt.Println(structlog.LogStruct(res))

// 	fmt.Println("================================================================")
// }
