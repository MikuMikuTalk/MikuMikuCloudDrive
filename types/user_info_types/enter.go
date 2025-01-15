package userinfo_types

type UserInfoRequest struct {
	Token string
}

type UserInfoResponse struct {
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}
