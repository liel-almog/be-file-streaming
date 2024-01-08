package models

type StartUploadDTO struct {
	FileName string `json:"fileName" binding:"required"`
	Size     int64  `json:"size" binding:"required"`
}
