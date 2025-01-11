package pwd

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logrus.Error("加密密码失败")
		return "", errors.New("加密密码失败")
	}
	return string(hashPwd), nil
}
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}
