package repositories_implementation

import (
	"context"
	"fmt"
	"io"
	"net/http"

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
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}
	defer client.Close()

	resp, err := http.Get(image)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read image data: %v", err)
	}

	objectName := ulid.Make().String()

	bucket := client.Bucket(c.BucketName)

	writer := bucket.Object(objectName).NewWriter(ctx)

	writer.ContentType = http.DetectContentType(imageData)

	if _, err := writer.Write(imageData); err != nil {
		writer.Close()
		return "", fmt.Errorf("failed to upload image to bucket: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to finalize upload to bucket: %v", err)
	}

	return objectName, nil
}
