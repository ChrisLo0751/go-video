package router

import (
	"go-video/base/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup, userHandler *api.UserHandler) {
	userRouter := router.Group("users")
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.POST("", userHandler.CreateUser)
}
