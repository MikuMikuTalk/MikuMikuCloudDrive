package user_service

import (
	"MikuMikuCloudDrive/models/user_models"
	"MikuMikuCloudDrive/types/resgister_types"
	"MikuMikuCloudDrive/utils/pwd"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
