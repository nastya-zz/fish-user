package user

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"user/internal/service"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService          service.UserService
	settingsService      service.SettingsService
	subscriptionsService service.SubscriptionsService
}

func NewImplementation(
	authService service.UserService,
	settingsService service.SettingsService,
	subscriptionsService service.SubscriptionsService,
) *Implementation {
	return &Implementation{
		authService:          authService,
		settingsService:      settingsService,
		subscriptionsService: subscriptionsService,
	}
}
