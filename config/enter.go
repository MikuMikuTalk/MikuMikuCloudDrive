package config

type Config struct {
	App         App                `mapstructure:"app"`
	MysqlConfig MySQLConfiguration `mapstructure:"mysql"`
}
type App struct {
	Title   string `mapstructure:"title"`
	Version string `mapstructure:"version"`
	Server  string `mapstructure:"server"`
	Port    uint   `mapstructure:"port"`
}
type MySQLConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
