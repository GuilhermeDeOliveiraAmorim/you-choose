package repositories_implementation

import (
	"context"
	"io"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/oklog/ulid/v2"

	"cloud.google.com/go/storage"
)

type ImageRepository struct {
	BucketName string
}

func NewImageRepository(bucketName string) *ImageRepository {
	return &ImageRepository{
		BucketName: bucketName,
	}
}

func (c *ImageRepository) SaveImage(image string) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := http.Get(image)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	objectName := ulid.Make().String()

	bucket := client.Bucket(c.BucketName)

	writer := bucket.Object(objectName).NewWriter(ctx)

	writer.ContentType = http.DetectContentType(imageData)

	if _, err := writer.Write(imageData); err != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: err.Error(),
			From:    "SaveImage",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		writer.Close()
		return "", err
	}

	if err := writer.Close(); err != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: err.Error(),
			From:    "SaveImage",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		return "", err
	}

	return objectName, nil
}
