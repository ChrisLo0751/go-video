package model

type File struct {
	ID       uint   `json:"id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
