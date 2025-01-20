package userinfo_types

type UserInfoRequest struct {
}

type UserInfoResponse struct {
	UserID   uint   `json:"user_id" testlog:"用户ID"`
	UserName string `json:"username" testlog:"用户名"`
	Avatar   string `json:"avatar" testlog:"头像"`
	Email    string `json:"email" testlog:"邮箱"`
}

type UpdateUserInfoRequest struct {
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInfoResponse struct{}
