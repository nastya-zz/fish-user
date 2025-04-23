package subscribtions

import (
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

func (s serv) Subscriptions(id model.UserId) *model.Subscriptions {
	//TODO implement me
	panic("implement me")
}
