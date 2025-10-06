package user

import (
	"context"
	"encoding/json"
	"fmt"

	"user/internal/model"
	"user/pkg/logger"
)

func (s serv) UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error) {
	const op = "user.UpdateProfile"

	var err error
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		updatedProfile, errTx := s.userRepository.UpdateProfile(ctx, updateInfo)
		if errTx != nil {
			return errTx
		}

		body, errTx := json.Marshal(model.UpdateProfile{
			ID:         updatedProfile.ID,
			Name:       updatedProfile.Name,
			AvatarPath: updatedProfile.AvatarPath,
			Bio:        updatedProfile.Bio,
			IsPublic:   updatedProfile.IsPublic,
		})

		if errTx != nil {
			return fmt.Errorf("error in marshal json body %w", errTx)
		}

		event := &model.Event{
			Type:    model.UserUpdateProfile,
			Payload: body,
		}

		errTx = s.eventRepository.SaveEvent(ctx, event)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		logger.Error(op, "error in update profile", err)
		return nil, fmt.Errorf("error in update profile %w", err)
	}

	return s.userRepository.UpdateProfile(ctx, updateInfo)
}

func (s serv) UpdateInfo(ctx context.Context, updateInfo *model.UpdateUser) error {
	//todo: add update in auth service by rabbit queue with outbox pattern

	return s.userRepository.UpdateInfo(ctx, updateInfo)
}
