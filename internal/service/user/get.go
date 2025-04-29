package user

import (
	"context"
	"user/internal/model"
)

func (s serv) UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error) {
	return s.userRepository.UserProfile(ctx, id)
}
