package user

import (
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"user/internal/service"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService          service.UserService
	settingsService      service.SettingsService
	subscriptionsService service.SubscriptionsService
}

func NewImplementation(
	userService service.UserService,
	settingsService service.SettingsService,
	subscriptionsService service.SubscriptionsService,
) *Implementation {
	return &Implementation{
		userService:          userService,
		settingsService:      settingsService,
		subscriptionsService: subscriptionsService,
	}
}
