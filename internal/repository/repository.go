package repository

import (
	"context"
	"user/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, profile *model.Profile) (model.UserId, error)
	UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error)
	UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error)
}

type SettingsRepository interface {
	Create(ctx context.Context, id model.UserId, settings model.Settings) (model.UserId, error)
	Get(ctx context.Context, id model.UserId) (*model.Settings, error)
	Update(ctx context.Context, id model.UserId, settings *model.Settings) (*model.Settings, error)
	Reset(ctx context.Context, id model.UserId) (*model.Settings, error)
}

type SubscriptionsRepository interface {
	Subscriptions(ctx context.Context, id model.UserId) *model.Subscriptions
}

type EventRepository interface {
	GetNewEvent(ctx context.Context) (*model.Event, error)
	SaveEvent(ctx context.Context, event *model.Event) error
	SetDone(ctx context.Context, id int) error
}
