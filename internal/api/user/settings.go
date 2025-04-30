package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/model"
)

func (i *Implementation) GetSettings(ctx context.Context, req *desc.GetSettingsRequest) (*desc.GetSettingsResponse, error) {
	id := req.GetId()
	if len(id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing user id")
	}

	settings, err := i.settingsService.Get(ctx, model.UserId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user settings")
	}

	return &desc.GetSettingsResponse{
		Settings: &desc.AccountSettings{
			Availability: getAvailability(settings.Availability),
			Language:     getLanguage(settings.Language),
		},
	}, nil
}

func getAvailability(str string) desc.Availability {
	switch str {
	case model.Private:
		return desc.Availability_PRIVATE
	case model.Public:
		return desc.Availability_PUBLIC
	default:
		return desc.Availability_PUBLIC
	}
}

func getLanguage(str string) desc.Language {
	switch str {
	case model.LangRu:
		return desc.Language_RU
	case model.LangEn:
		return desc.Language_ENG
	default:
		return desc.Language_RU
	}
}
