package user

import (
	"context"
	"user/internal/model"
)

func (s serv) DeleteUser(ctx context.Context, id model.UserId) error {
	return s.userRepository.DeleteUser(ctx, id)
}
