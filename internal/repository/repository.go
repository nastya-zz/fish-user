package repository

import (
	"context"
	"user/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, profile *model.Profile) error
	UserProfile(ctx context.Context, id model.UserId) *model.Profile
	UpdateProfile(ctx context.Context, id model.UserId, updateInfo *model.UpdateProfile) *model.Profile
}

type SettingsRepository interface {
	Settings(ctx context.Context, id model.UserId) *model.Settings
}

type SubscriptionsRepository interface {
	Subscriptions(ctx context.Context, id model.UserId) *model.Subscriptions
}
