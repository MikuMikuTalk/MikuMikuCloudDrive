package user_service

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/models/user_models"
	userinfo_types "MikuMikuCloudDrive/types/user_info_types"
	"MikuMikuCloudDrive/utils/jwts"
	"errors"

	"github.com/sirupsen/logrus"
)

func (s *UserService) GetUserInfo(getUserInfoReq userinfo_types.UserInfoRequest) (*userinfo_types.UserInfoResponse, error) {
	token := jwts.ProcessJwtToken(getUserInfoReq.Token)
	authConfig := config.ReadAuthConfig()
	claims, err := jwts.ParseJwtToken(token, authConfig.AuthSecret)
	if err != nil {
		logrus.Error("jwt 解析失败")
		return nil, errors.New("jwt 解析失败")
	}

	var userModel user_models.UserModel
	err = s.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		logrus.Error("查找用户失败")
		return nil, errors.New("查找用户失败")
	}
	return &userinfo_types.UserInfoResponse{
		UserName: userModel.UserName,
		Email:    userModel.Email,
		Avatar:   userModel.Avatar,
	}, nil
}

func (s *UserService) UpdateUserInfo(updateUserInfoReq userinfo_types.UpdateUserInfoRequest) (*userinfo_types.UpdateUserInfoResponse, error) {
	token := jwts.ProcessJwtToken(updateUserInfoReq.Token)
	authConfig := config.ReadAuthConfig()
	claims, err := jwts.ParseJwtToken(token, authConfig.AuthSecret)
	if err != nil {
		logrus.Error("jwt 解析失败")
		return nil, errors.New("jwt 解析失败")
	}
	userId := claims.UserID
	var userModel user_models.UserModel
	err = s.DB.Take(&userModel, userId).Error
	if err != nil {
		logrus.Error("用户不存在")
		return nil, errors.New("用户不存在")
	}
	if updateUserInfoReq.UserName != "" {
		userModel.UserName = updateUserInfoReq.UserName
	}
	if updateUserInfoReq.Avatar != "" {
		userModel.Avatar = updateUserInfoReq.Avatar
	}
	if updateUserInfoReq.Email != "" {
		userModel.Email = updateUserInfoReq.Avatar
	}
	if updateUserInfoReq.Password != "" {
		userModel.Password = updateUserInfoReq.Password
	}
	err = s.DB.Save(&userModel).Error
	if err != nil {
		logrus.Error("用户信息更新失败")
		return nil, errors.New("用户信息更新失败")
	}
	return &userinfo_types.UpdateUserInfoResponse{}, nil
}
