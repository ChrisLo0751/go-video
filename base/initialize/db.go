package initialize

import (
	"fmt"
	"go-video/base/config"
	"go-video/base/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := GetDBConfig(cfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Optional: 自动迁移模式
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserInfo{})

	return db, nil
}

func GetDBConfig(cfg config.DatabaseConfig) string {
	host := cfg.Host
	port := cfg.Port
	username := cfg.User
	password := cfg.Password
	dbname := cfg.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	return dsn
}
