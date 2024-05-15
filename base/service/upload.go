package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type FileService struct {
}

type FileImpl interface {
	UploadFiles(files []*multipart.FileHeader) error
}

func NewFileService() FileImpl {
	return &FileService{}
}

func (f FileService) UploadFiles(files []*multipart.FileHeader) error {
	for _, file := range files {
		dst, err := os.Create("uploads" + file.Filename)
		if err != nil {
			return fmt.Errorf("failed to create destination file %s", file.Filename)
		}
		defer dst.Close()

		src, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open uploaded file %s", file.Filename)
		}
		defer src.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			return fmt.Errorf("failed to save uploaded file %s", file.Filename)
		}
	}

	return nil
}
