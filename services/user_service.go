package services

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models/user_models"
	"MikuMikuCloudDrive/types/login_types"
	"MikuMikuCloudDrive/types/resgister_types"
	"MikuMikuCloudDrive/utils/jwts"
	"MikuMikuCloudDrive/utils/pwd"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}
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
func (us *UserService) Register(username, password string) (*resgister_types.RegisterResponse, error) {
	err := us.DB.Take(&user_models.UserModel{}, "user_name = ?", username).Error
	if err != nil {
		logrus.Infof("%s 不存在，将要创建用户", username)
	}
	encryptPassword, err := pwd.EncryptPassword(password)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var user user_models.UserModel = user_models.UserModel{
		UserName: username,
		Password: encryptPassword,
		Email:    username + "@example.com",
		Avatar:   "https://codeberg.org/avatars/ce7f1e613bcdcdc206ad744919743ba556d33d6cd30de1a64e0ddb119de4c8c0?size=512",
	}
	err = us.DB.Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err1 := errors.New("用户已经注册过了，不要重复注册")
			errorMessage := fmt.Sprintf("创建用户%s失败%v", username, err.Error())
			logrus.Error(errorMessage)
			return nil, err1
		}
		return nil, err
	}
	logrus.Info("用户创建成功")
	return nil, err
}
