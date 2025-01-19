package userinfo_types

type UserInfoRequest struct {
	Token string
}

type UserInfoResponse struct {
	UserName string `json:"username" testlog:"用户名"`
	Avatar   string `json:"avatar" testlog:"头像"`
	Email    string `json:"email" testlog:"邮箱"`
}

type UpdateUserInfoRequest struct {
	Token    string `header:"Authorization"` // 绑定到 Authorization 头
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInfoResponse struct{}
