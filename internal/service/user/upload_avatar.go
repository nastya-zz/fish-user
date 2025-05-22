package user

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"path/filepath"
)

func (s serv) UploadAvatar(ctx context.Context, file []byte, name string) (string, error) {
	const op = "service.User.UploadAvatar"

	ext := filepath.Ext(name)
	contentType, err := getContentTYpe(ext)
	if err != nil {
		return "", err
	}

	nameId := uuid.NewString()
	uniqUuidName := fmt.Sprintf("%s%s", nameId, ext)

	ui, err := s.minio.Client.PutObject(ctx, s.minio.BucketName, uniqUuidName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("%s: uploaded file %s to bucket %s\n", op, ui.ETag, s.minio.BucketName)

	srcLink := s.minio.SourcePrefix()
	link := fmt.Sprintf("%s%s", srcLink, uniqUuidName)

	return link, nil
}

func getContentTYpe(ext string) (string, error) {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg", nil
	case ".png":
		return "image/png", nil
	case ".webp":
		return "image/webp", nil
	case ".heic":
		return "image/heic", nil

	default:
		return "", fmt.Errorf("extention not support")
	}
}
