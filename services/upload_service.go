package services

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/lielalmog/be-file-streaming/models"
	"github.com/lielalmog/be-file-streaming/repositories"
)

type UploadService interface {
	StartUpload(fileMetadate *models.FileMetadateDTO) (*int64, error)
	UploadChunk(fileHeader *multipart.FileHeader, id int64, chunkIndex int) error
}

type uploadServiceImpl struct{}

var (
	initUploadServiceOnce sync.Once
	uploadService         *uploadServiceImpl
)

const CONTAINER_NAME = "files"

func (u *uploadServiceImpl) StartUpload(fileMetadate *models.FileMetadateDTO) (*int64, error) {
	return repositories.GetUploadRepository().SaveFileMetadata(fileMetadate)
}

func (u *uploadServiceImpl) UploadChunk(fileHeader *multipart.FileHeader, id int64, chunkIndex int) error {
	connectionString, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		log.Fatal("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	serviceClient, err := azblob.NewClientFromConnectionString(connectionString, nil)

	blobName := fmt.Sprintf("%d/%d", id, chunkIndex)
	fmt.Println(blobName)
	serviceClient.UploadStream(context.Background(), CONTAINER_NAME, blobName, file, nil)
	if err != nil {
		return err
	}

	// file, err := fileHeader.Open()

	// if err != nil {
	// 	return err
	// }
	// fmt.Println(file)

	return nil
}

func newUploadService() *uploadServiceImpl {
	return &uploadServiceImpl{}
}

func GetUploadService() UploadService {
	initUploadServiceOnce.Do(func() {
		uploadService = newUploadService()
	})

	return uploadService
}
