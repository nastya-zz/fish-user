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

func (i *Implementation) ResetSettings(ctx context.Context, req *desc.ResetSettingsRequest) (*desc.ResetSettingsResponse, error) {
	userId := req.GetUserId()
	if len(userId) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}
	updated, err := i.settingsService.Reset(ctx, model.UserId(userId))
	if err != nil {
		return nil, status.Error(codes.Internal, api_errors.UserUpdateSettingsFailed)
	}

	return &desc.ResetSettingsResponse{
		Settings: &desc.AccountSettings{
			Availability: converter.GetDescAvailability(updated.Availability),
			Language:     converter.GetDescLanguage(updated.Language),
		},
	}, nil

}
