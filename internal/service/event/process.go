package event

import (
	"context"
	"errors"
	"fmt"
	"user/internal/converter"
	"user/internal/logger"
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
		return err
	case model.UserUpdate:
		logger.Info(op, "event", model.UserUpdate, "payload", event.Payload)
		return fmt.Errorf("update is not implement %w", ErrNotImplementedEventType)
	default:
		return fmt.Errorf("can't process message %w", ErrUnknownEventType)
	}
}
