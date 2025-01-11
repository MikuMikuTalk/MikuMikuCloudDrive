package core

import (
	"MikuMikuCloudDrive/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	var mysqlConfig config.MySQLConfiguration
	mysqlViper := viper.New()

	mysqlViper.AddConfigPath("./config")
	mysqlViper.SetConfigName("config")
	mysqlViper.SetConfigType("toml")
	if err := mysqlViper.ReadInConfig(); err != nil {
		logrus.Fatalf("read config failed: %v", err)
	}
	// 仅解析 `mysql` 部分到结构体中
	if err := mysqlViper.Sub("mysql").Unmarshal(&mysqlConfig); err != nil {
		logrus.Fatalf("unmarshal mysql config failed: %v", err)
	}

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
