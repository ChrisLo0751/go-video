package api

import (
	"github.com/gin-gonic/gin"
	"go-video/base/service"
	"net/http"
)

type FileHandler struct {
	fileService service.FileImpl
}

func NewFileHandler(fileService service.FileImpl) *FileHandler {
	return &FileHandler{
		fileService: fileService,
	}
}

func (f *FileHandler) UploadFiles(c *gin.Context) {
	files := c.Request.MultipartForm.File["files"]

	err := f.fileService.UploadFiles(files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to uploaded files",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Files uploaded and saved successfully",
	})
}
