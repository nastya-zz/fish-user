package subscriptions

import (
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SubscriptionsRepository {
	return &repo{db: db}
}

func (r repo) Subscriptions(id model.UserId) *model.Subscriptions {
	//TODO implement me
	panic("implement me")
}
