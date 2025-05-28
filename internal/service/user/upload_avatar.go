package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
	"user/internal/logger"
	"user/internal/model"
)

func (s serv) UploadAvatar(ctx context.Context, file []byte, name string, userId model.UserId) (string, error) {
	const op = "service.User.UploadAvatar"

	ext := filepath.Ext(name)

	nameId := uuid.NewString()
	uniqUuidName := fmt.Sprintf("%s%s", nameId, ext)

	link, err := s.minio.UploadFile(ctx, file, uniqUuidName)

	oldLink, err := s.oldAvatarLink(ctx, userId)
	if err != nil {
		return "", err
	}

	if len(oldLink) > 0 {
		removeAvatarErr := s.minio.RemoveFile(ctx, oldLink)
		if removeAvatarErr != nil {
			fmt.Printf("%s: %s", op, removeAvatarErr.Error())
		}
	}

	return link, nil
}

func (s serv) oldAvatarLink(ctx context.Context, userId model.UserId) (string, error) {
	const op = "service.User.oldAvatarLink"

	userProfile, err := s.UserProfile(ctx, userId)
	if err != nil || userProfile == nil {
		return "", fmt.Errorf("cannot chexk user profile %s, %w", op, err)
	}

	logger.Info(op, "userProfile.AvatarPath old link: ", userProfile.AvatarPath)
	return userProfile.AvatarPath, nil
}
