package go_test

// func TestCreateDirectory(t *testing.T) {
// 	db := core.InitGorm()
// 	// rdb := core.InitRedis()

// 	directory_svc := directory_service.NewDirectoryService(db)
// 	req := directory_types.CreateDirectoryRequest{
// 		Name: "meowrain",
// 	}
// 	res, err := directory_svc.CreateDirectory(req, &jwts.CustomClaims{
// 		UserID:   1,
// 		Username: "meowrain",
// 	})
// 	if err != nil {
// 		t.Log(err)
// 		return
// 	}
// 	fmt.Println(res)
// }
