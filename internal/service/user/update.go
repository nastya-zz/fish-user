package user

import (
	"context"
	"user/internal/model"
)

func (s serv) UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error) {
	//todo: add update in auth service by rabbit queue with outbox pattern

	return s.userRepository.UpdateProfile(ctx, updateInfo)
}
