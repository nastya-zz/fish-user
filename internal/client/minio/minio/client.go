package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client     *minio.Client
	bucketName string
}

const bucketName = "user-service"

func New(ctx context.Context, endpoint, accessKeyID, secretAccessKey string) (*Client, error) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, fmt.Errorf("произошла ошибка при  создании клиента минио %w", err)
	}

	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil && !exists {
		return nil, fmt.Errorf("произошла ошибка при  проверке бакета %w", err)
	}

	return &Client{
		client:     minioClient,
		bucketName: bucketName,
	}, err
}
