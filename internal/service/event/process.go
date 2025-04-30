package event

import (
	"context"
	"errors"
	"fmt"
	"user/internal/converter"
	"user/internal/model"
)

var ErrUnknownEventType = errors.New("Process.UnknownEventType")

func (p Processor) Process(ctx context.Context, event model.Event) error {
	switch event.Type {
	case model.UserCreate:
		user := converter.UserFromPayload(event.Payload)
		_, err := p.userService.SaveUser(ctx, &user)
		return err
	default:
		return fmt.Errorf("can't process message %w", ErrUnknownEventType)
	}
}
