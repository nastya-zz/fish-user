package subscribtions

import (
	"context"
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/service"
)

type serv struct {
	subscriptionsRepository repository.SubscriptionsRepository
	txManager               db.TxManager
}

func NewService(
	subscriptionsRepository repository.SubscriptionsRepository,
	txManager db.TxManager,
) service.SubscriptionsService {
	return &serv{
		subscriptionsRepository: subscriptionsRepository,
		txManager:               txManager,
	}
}

func (s serv) Subscriptions(ctx context.Context, id model.UserId) (*model.Subscriptions, error) {
	return s.subscriptionsRepository.Subscriptions(ctx, id)
}

func (s serv) Subscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	return s.subscriptionsRepository.Subscribe(ctx, id, subscriptionId)
}

func (s serv) Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	return s.subscriptionsRepository.Unsubscribe(ctx, id, subscriptionId)
}
