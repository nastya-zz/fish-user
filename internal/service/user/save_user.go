package user

import (
	"context"
	"fmt"
	"user/internal/model"
)

func (s serv) SaveUser(ctx context.Context, profile *model.Profile) (model.UserId, error) {
	const op = "user.SaveUser"

	var id model.UserId
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.SaveUser(ctx, profile)
		if errTx != nil {
			return fmt.Errorf("%s %w", op, errTx)
		}

		_, errTx = s.settingsService.Create(ctx, id)
		if errTx != nil {
			return fmt.Errorf("%s %w", op, errTx)
		}

		return nil
	})

	return id, err
}
