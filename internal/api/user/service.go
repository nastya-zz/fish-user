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

//func (UnimplementedUserV1Server) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
//}
//func (UnimplementedUserV1Server) UpdateSettings(context.Context, *UpdateSettingsRequest) (*UpdateSettingsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateSettings not implemented")
//}
//func (UnimplementedUserV1Server) ResetSettings(context.Context, *ResetSettingsRequest) (*ResetSettingsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ResetSettings not implemented")
//}
//func (UnimplementedUserV1Server) GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
