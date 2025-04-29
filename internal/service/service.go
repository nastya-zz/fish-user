package service

import (
	"context"
	"user/internal/model"
)

type UserService interface {
	SaveUser(context.Context, *model.Profile) error
	UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error)
	UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error)
}

type SettingsService interface {
	Settings(id model.UserId) *model.Settings
}

type SubscriptionsService interface {
	Subscriptions(id model.UserId) *model.Subscriptions
}

type EventsService interface {
	Process(ctx context.Context, event model.Event) error
}
