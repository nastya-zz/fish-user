package service

import (
	"context"
	"user/internal/model"
)

type UserService interface {
	SaveUser(context.Context, *model.Profile) error
	UserProfile(id model.UserId) *model.Profile
	UpdateProfile(id model.UserId, updateInfo *model.UpdateProfile) *model.Profile
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
