package service

import (
	"context"
	"user/internal/model"
)

type UserService interface {
	SaveUser(context.Context, *model.Profile) (model.UserId, error)
	UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error)
	UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error)
}

type SettingsService interface {
	Create(ctx context.Context, id model.UserId) (model.UserId, error)
	Get(ctx context.Context, id model.UserId) (*model.Settings, error)
	Update(ctx context.Context, id model.UserId, settings *model.Settings) (*model.Settings, error)
	Reset(ctx context.Context, id model.UserId) (*model.Settings, error)
}

type SubscriptionsService interface {
	Subscriptions(ctx context.Context, id model.UserId) (*model.Subscriptions, error)
	Subscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error
	Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error
}

type EventsService interface {
	Process(ctx context.Context, event model.Event) error
}
