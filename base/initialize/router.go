package initialize

import (
	"go-video/user-web/api"
	"go-video/user-web/config"
	"go-video/user-web/data"
	"go-video/user-web/router"
	"go-video/user-web/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routers(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()

	userData := data.NewUserData(db)
	userService := service.NewUserService(userData, cfg.Auth.JWTSecret)
	userHandler := api.NewUserHandler(userService)

	group := r.Group("/v1")
	router.InitUserRouter(group, userHandler)

	return r
}
