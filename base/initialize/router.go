package initialize

import (
	"go-video/base/api"
	"go-video/base/config"
	"go-video/base/data"
	"go-video/base/router"
	"go-video/base/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routers(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()
	group := r.Group("/v1")

	// 用户的api
	userData := data.NewUserData(db)
	userService := service.NewUserService(userData, cfg.Auth.JWTSecret)
	userHandler := api.NewUserHandler(userService)
	router.InitUserRouter(group, userHandler)

	// 文件的api
	fileService := service.NewFileService()
	fileHandler := api.NewFileHandler(fileService)
	router.InitFileRouter(group, fileHandler)

	return r
}
