package event

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"user/internal/client/broker/rabbitmq"
	"user/internal/model"
	"user/pkg/logger"
)

const (
	eventBatchSize = 20
)

func (s eventService) StartPublishEvents(ctx context.Context, handlePeriod time.Duration) {
	const op = "services.event-sender.StartProcessEvents"

	ticker := time.NewTicker(handlePeriod)

	go func() {
		for {
			select {
			case <-ctx.Done():
				logger.Info("stopping event processing", "op", op)
				return
			case <-ticker.C:
				// noop
			}

			events, err := s.sender.eventRepository.GetNewEvent(ctx, eventBatchSize)
			if err != nil {
				logger.Error("failed to get new event", "error", err, "op", op)
				continue
			}
			if events == nil {
				logger.Debug("no new events", "op", op)
				continue
			}

			s.SendMessage(ctx, events)
		}
	}()
}

func (s eventService) SendMessage(ctx context.Context, events []*model.Event) {
	const op = "services.event-sender.SendMessage"

	for _, event := range events {
		logger.Info("sending message", "event", event, "op", op)

		err := s.publish(ctx, event)
		if err != nil {
			logger.Error("failed to send message", "error", err, "op", op)
		}

		if err := s.sender.eventRepository.SetDone(ctx, event.ID); err != nil {
			logger.Error("failed to set event done", "error", err, "op", op)
		}
	}

}

func (s eventService) SendEvent(ctx context.Context, event *model.Event) error {
	return s.publish(ctx, event)
}

func (s eventService) publish(_ context.Context, event *model.Event) error {
	const op = "broker.publish"
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(event); err != nil {
		return fmt.Errorf("could not encode event: %w", err)
	}
	err := s.sender.ch.Publish(
		rabbitmq.ExchangeName, // exchange
		"",                    // routing key (не используется для fanout)
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			AppId:       "auth_grpc_server",
			ContentType: "application/x-encoding-gob",
			Body:        b.Bytes(),
			Timestamp:   time.Now(),
		})
	if err != nil {
		return fmt.Errorf("could not publish: %s, %w", op, err)
	}

	return nil
}
