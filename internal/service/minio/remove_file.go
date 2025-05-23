package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"strings"
)

func (s Service) RemoveFile(ctx context.Context, link string) error {
	fileName, err := s.getFileName(link)
	if err != nil {
		return err
	}

	err = s.minioClient.Client.RemoveObject(ctx, s.minioClient.BucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("не удалось удалить объект: %v", err)
	}

	log.Printf("Файл %s успешно удален из бакета, имя файла: %s", s.minioClient.BucketName, fileName)
	return nil
}

func (s Service) getFileName(link string) (string, error) {
	name, found := strings.CutPrefix(link, s.minioClient.BucketName+"/")
	if !found {
		return "", fmt.Errorf("cannot get filename")
	}

	return name, nil
}
