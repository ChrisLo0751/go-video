package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type AuthConfig struct {
	JWTSecret  string
	PublicKey  string
	PrivateKey string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // 读取匹配的环境变量

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
