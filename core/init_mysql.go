package core

import (
	"MikuMikuCloudDrive/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm() *gorm.DB {
	var mysqlConfig config.MySQLConfiguration = config.ReadMySQLConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.Error("连接Mysql失败")
		panic(err)
	} else {
		logrus.Debug("连接数据库成功")
	}
	return db
}
