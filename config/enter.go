package config

import "github.com/spf13/viper"

type Config struct {
	App         App                `mapstructure:"app"`
	MysqlConfig MySQLConfiguration `mapstructure:"mysql"`
}
type App struct {
	Title     string `mapstructure:"title"`
	Version   string `mapstructure:"version"`
	Server    string `mapstructure:"server"`
	Port      uint   `mapstructure:"port"`
	UploadDir string `mapstructure:"upload_dir"`
	CacheDir  string `mapstructure:"cache_dir"`
}
type MySQLConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type RedisConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"db"`
}
type Auth struct {
	AuthSecret string `mapstructure:"auth_secret"`
	ExpireTime int    `mapstructure:"expire_time"`
}

func ReadAppConfig() App {
	var appConfiguration App
	mainConfiguration := viper.New()
	mainConfiguration.SetConfigName("config")
	mainConfiguration.AddConfigPath("./config")
	mainConfiguration.SetConfigType("toml")
	if err := mainConfiguration.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := mainConfiguration.Sub("app").Unmarshal(&appConfiguration); err != nil {
		panic(err)
	}
	return appConfiguration
}
func ReadMySQLConfig() MySQLConfiguration {
	var mysqlConfiguration MySQLConfiguration
	mainConfiguration := viper.New()
	mainConfiguration.SetConfigName("config")
	mainConfiguration.AddConfigPath("./config")
	mainConfiguration.SetConfigType("toml")
	if err := mainConfiguration.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := mainConfiguration.Sub("mysql").Unmarshal(&mysqlConfiguration); err != nil {
		panic(err)
	}
	return mysqlConfiguration
}
func ReadRedisConfig() RedisConfiguration {
	var redisConfiguration RedisConfiguration
	mainConfiguration := viper.New()
	mainConfiguration.SetConfigName("config")
	mainConfiguration.AddConfigPath("./config")
	mainConfiguration.SetConfigType("toml")
	if err := mainConfiguration.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := mainConfiguration.Sub("redis").Unmarshal(&redisConfiguration); err != nil {
		panic(err)
	}
	return redisConfiguration
}
func ReadAuthConfig() Auth {
	var authConfiguration Auth
	mainConfiguration := viper.New()
	mainConfiguration.SetConfigName("config")
	mainConfiguration.AddConfigPath("./config")
	mainConfiguration.SetConfigType("toml")
	if err := mainConfiguration.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := mainConfiguration.Sub("auth").Unmarshal(&authConfiguration); err != nil {
		panic(err)
	}
	return authConfiguration
}
