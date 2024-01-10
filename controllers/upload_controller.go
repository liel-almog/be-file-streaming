package controllers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type UploadController interface {
	StartUpload(c *fiber.App) error
}

type UploadControllerImpl struct{}

func (u *UploadControllerImpl) StartUpload(c *fiber.App) error {
	return nil
}

var (
	initUploadControllerOnce sync.Once
	uploadController         *UploadControllerImpl
)

func NewUploadController() *UploadControllerImpl {
	return &UploadControllerImpl{}
}

func GetUploadController() UploadController {
	initUploadControllerOnce.Do(func() {
		uploadController = NewUploadController()
	})

	return uploadController
}
