package subscriptions

import (
	"context"
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

func (r repo) Subscriptions(ctx context.Context, id model.UserId) *model.Subscriptions {
	//TODO implement me
	panic("implement me")
}
