package resgister_types

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterResponse struct{}
