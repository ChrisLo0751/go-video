package router

import (
	"github.com/gin-gonic/gin"
	"go-video/base/api"
)

func InitFileRouter(router *gin.RouterGroup, handler *api.FileHandler) {
	userRouter := router.Group("files")
	userRouter.POST("", handler.UploadFiles)
}
