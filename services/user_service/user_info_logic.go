package user_service

import (
	"errors"

	"MikuMikuCloudDrive/models"

	userinfo_types "MikuMikuCloudDrive/types/user_info_types"
	"MikuMikuCloudDrive/utils/jwts"

	"github.com/sirupsen/logrus"
)

func (s *UserService) GetUserInfo(getUserInfoReq userinfo_types.UserInfoRequest, claims *jwts.CustomClaims) (*userinfo_types.UserInfoResponse, error) {

	var userModel models.UserModel
	err := s.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		logrus.Error("查找用户失败")
		return nil, errors.New("查找用户失败")
	}
	return &userinfo_types.UserInfoResponse{
		UserID:   userModel.ID,
		UserName: userModel.UserName,
		Email:    userModel.Email,
		Avatar:   userModel.Avatar,
	}, nil
}

func (s *UserService) UpdateUserInfo(updateUserInfoReq userinfo_types.UpdateUserInfoRequest, claims *jwts.CustomClaims) (*userinfo_types.UpdateUserInfoResponse, error) {

	userId := claims.UserID
	var userModel models.UserModel
	err := s.DB.Take(&userModel, userId).Error
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
