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

//func (UnimplementedUserV1Server) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
//}
//func (UnimplementedUserV1Server) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
//}
//func (UnimplementedUserV1Server) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
//}
//func (UnimplementedUserV1Server) GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
//}
