package user

import (
	"context"
	"user/internal/model"
)

func (s serv) UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error) {
	return s.userRepository.UpdateProfile(ctx, updateInfo)
}
