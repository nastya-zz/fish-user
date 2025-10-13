package subscribtions

import (
	"context"
	"encoding/json"
	"fmt"

	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/logger"
)

type serv struct {
	subscriptionsRepository repository.SubscriptionsRepository
	txManager               db.TxManager
	eventRepository         repository.EventRepository
}

func NewService(
	subscriptionsRepository repository.SubscriptionsRepository,
	txManager db.TxManager,
	eventRepository repository.EventRepository,
) service.SubscriptionsService {
	return &serv{
		subscriptionsRepository: subscriptionsRepository,
		txManager:               txManager,
		eventRepository:         eventRepository,
	}
}

func (s serv) Subscriptions(ctx context.Context, id model.UserId) (*model.Subscriptions, error) {
	return s.subscriptionsRepository.Subscriptions(ctx, id)
}

func (s serv) Subscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	const op = "user.Subscribe"

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		exists, err := s.subscriptionsRepository.SubscriptionExists(ctx, id, subscriptionId)
		if err != nil {
			return err
		}

		if exists {
			return fmt.Errorf("subscription already exists")
		}

		err = s.subscriptionsRepository.Subscribe(ctx, id, subscriptionId)
		if err != nil {
			logger.Error(op, "err", err)
			return err
		}

		err = s.createEvent(ctx, model.UserSubscribed, model.SubscribePayload{
			FollowerID: id.String(),
			AuthorID:   subscriptionId.String(),
		})
		if err != nil {
			logger.Error(op, "err", err)
			return err
		}

		return nil
	})

	return err
}

func (s serv) Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	const op = "user.Unsubscribe"

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := s.subscriptionsRepository.Unsubscribe(ctx, id, subscriptionId)
		if err != nil {
			logger.Error(op, "err", err)
			return err
		}

		return s.createEvent(ctx, model.UserUnsubscribed, model.SubscribePayload{
			FollowerID: id.String(),
			AuthorID:   subscriptionId.String(),
		})
	})

	return err
}

func (s serv) createEvent(ctx context.Context, eventType string, payload model.SubscribePayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	event := &model.Event{
		Type:    eventType,
		Payload: body,
	}

	return s.eventRepository.SaveEvent(ctx, event)
}
