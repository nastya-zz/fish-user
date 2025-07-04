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

func (i *Implementation) UpdateSettings(ctx context.Context, req *desc.UpdateSettingsRequest) (*desc.UpdateSettingsResponse, error) {
	userId := req.GetUserId()
	if len(userId) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	setInfo := mappingSettings(req.GetSettingsInfo())

	updated, err := i.settingsService.Update(ctx, model.UserId(userId), setInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, api_errors.UserUpdateSettingsFailed)
	}

	return &desc.UpdateSettingsResponse{
		Settings: &desc.AccountSettings{
			Availability: converter.GetDescAvailability(updated.Availability),
			Language:     converter.GetDescLanguage(updated.Language),
		},
	}, nil
}
func mappingSettings(s *desc.AccountSettings) *model.Settings {
	return &model.Settings{
		Language:     converter.GetModelLanguage(s.Language),
		Availability: converter.GetModelAvailability(s.Availability),
	}
}
