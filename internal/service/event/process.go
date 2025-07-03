package event

import (
	"context"
	"errors"
	"fmt"
	"user/internal/converter"
	"user/pkg/logger"
	"user/internal/model"
)

var ErrUnknownEventType = errors.New("Process.UnknownEventType")
var ErrNotImplementedEventType = errors.New("Process.NotImplementedEventType")

//todo добавить обработку апдейта и удалпения пользователя

func (p Processor) Process(ctx context.Context, event model.Event) error {
	const op = "process.Process"
	switch event.Type {

	case model.UserCreate:
		user := converter.UserFromPayload(event.Payload)
		_, err := p.userService.SaveUser(ctx, &user)
		if err != nil {
			logger.Error(op, "event", model.UserCreate, "err", err)
		}

		return err

	case model.UserUpdate:
		userForUpdate := converter.UpdateUserFromPayload(event.Payload)
		logger.Info(op, "event", model.UserUpdate, "payload", event.Payload, "userForUpdate", userForUpdate)

		err := p.userService.UpdateInfo(ctx, &userForUpdate)
		if err != nil {
			logger.Error(op, "event", model.UserUpdate, "err", err)
		}

	case model.UserDelete:
		user := converter.UserFromPayload(event.Payload)
		logger.Info(op, "event", model.UserUpdate, "payload", event.Payload, "user", user)

		err := p.userService.DeleteUser(ctx, user.ID)
		if err != nil {
			logger.Error(op, "event", model.UserUpdate, "err", err)
		}

		return err
	}

	return fmt.Errorf("can't process message %w", ErrUnknownEventType)
}
