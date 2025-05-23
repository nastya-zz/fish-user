package user

import (
	"context"
	"fmt"
	"log"
	"user/internal/model"
)

func (s serv) RemoveAvatar(ctx context.Context, userId model.UserId, filename string) error {
	const op = "service.User.RemoveAvatar"

	userProfile, err := s.UserProfile(ctx, userId)
	if err != nil || userProfile == nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if userProfile.AvatarPath != filename {
		return fmt.Errorf("user avatar link not match %s: %w", op, err)
	}

	updateProfile := &model.UpdateProfile{
		ID:         userId,
		AvatarPath: "",
	}

	_, err = s.userRepository.UpdateProfile(ctx, updateProfile)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = s.minio.RemoveFile(ctx, filename); err != nil {
		log.Print("minio remove file error:", err.Error())
	}

	return nil
}
