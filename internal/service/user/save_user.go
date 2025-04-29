package user

import (
	"context"
	"user/internal/model"
)

func (s serv) SaveUser(ctx context.Context, profile *model.Profile) error {
	return s.userRepository.SaveUser(ctx, profile)
}
