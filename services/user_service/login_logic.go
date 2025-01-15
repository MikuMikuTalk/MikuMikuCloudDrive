package user_service

import (
	"errors"
	"fmt"

	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models/user_models"
	"MikuMikuCloudDrive/types/login_types"
	"MikuMikuCloudDrive/utils/jwts"
	"MikuMikuCloudDrive/utils/pwd"

	"gorm.io/gorm"
)

func (us *UserService) Login(username, password string) (*login_types.LoginResponse, error) {
	var user user_models.UserModel
	err := us.DB.Take(&user, "user_name = ?", username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("数据库错误❌:%w", err)
	}
	payload := jwts.JwtPayload{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}
	authConfig := config.ReadAuthConfig()
	token, err := jwts.GenerateJwtToken(payload, authConfig.AuthSecret, authConfig.ExpireTime)
	if err != nil {
		return nil, errors.New("jwt 创建失败")
	}
	if !pwd.ComparePasswords(user.Password, password) {
		return nil, errors.New("用户密码错误")
	}
	return &login_types.LoginResponse{
		Token: token,
	}, nil
}
