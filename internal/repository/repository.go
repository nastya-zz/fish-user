package repository

import "user/internal/model"

type UserRepository interface {
	SaveUser(*model.Profile) error
	UserProfile(id model.UserId) *model.Profile
	UpdateProfile(id model.UserId, updateInfo *model.UpdateProfile) *model.Profile
}

type SettingsRepository interface {
	Settings(id model.UserId) *model.Settings
}

type SubscriptionsRepository interface {
	Subscriptions(id model.UserId) *model.Subscriptions
}
