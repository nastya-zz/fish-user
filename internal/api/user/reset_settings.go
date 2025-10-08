package user

import (
	"context"

	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user/internal/converter"
	"user/internal/model"
	"user/internal/utils"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) ResetSettings(ctx context.Context, req *desc.ResetSettingsRequest) (*desc.ResetSettingsResponse, error) {
	userID, err := utils.GetUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	updated, err := i.settingsService.Reset(ctx, model.UserId(userID.String()))
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
