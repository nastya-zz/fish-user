package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/converter"
	"user/internal/model"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) GetSettings(ctx context.Context, req *desc.GetSettingsRequest) (*desc.GetSettingsResponse, error) {
	id := req.GetId()
	if len(id) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	settings, err := i.settingsService.Get(ctx, model.UserId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, api_errors.UserGetSettingsFailed)
	}

	return &desc.GetSettingsResponse{
		Settings: &desc.AccountSettings{
			Availability: converter.GetDescAvailability(settings.Availability),
			Language:     converter.GetDescLanguage(settings.Language),
		},
	}, nil
}
