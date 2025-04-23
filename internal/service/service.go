package service

import "user/internal/model"

type UserService interface {
	SaveUser(*model.Profile) error
	UserProfile(id model.UserId) *model.Profile
	UpdateProfile(id model.UserId, updateInfo *model.UpdateProfile) *model.Profile
}

type SettingsService interface {
	Settings(id model.UserId) *model.Settings
}

type SubscriptionsService interface {
	Subscriptions(id model.UserId) *model.Subscriptions
}
