package user

import (
	"context"
	"strings"

	"github.com/google/uuid"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user/internal/converter"
	"user/internal/model"
	"user/internal/utils"
	api_errors "user/pkg/api-errors"
	"user/pkg/logger"
)

func (i *Implementation) UpdateProfile(ctx context.Context, req *desc.UpdateProfileRequest) (*desc.UpdateProfileResponse, error) {
	//todo убрать из обновления имя и емейл, надо получать их из сервис auth
	const op = "api.user.UpdateProfile"
	info := req.GetInfo()

	userID, err := utils.GetUserIdFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	if len(info.Email.Value) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserEmailEmpty)
	} else if checkEmailPattern(info.Email.Value) {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserEmailNotMatchPattern)
	}

	profile, err := i.userService.UpdateProfile(ctx, mappingUpdateProfile(info, userID))

	if err != nil {
		return nil, status.Error(codes.Internal, api_errors.UserUpdateFailed)
	}

	logger.Info(op, "updating user ", profile)

	return &desc.UpdateProfileResponse{Profile: converter.ToDescProfileFromProfile(profile)}, nil
}

func mappingUpdateProfile(info *desc.UpdateProfile, userID uuid.UUID) *model.UpdateProfile {
	return &model.UpdateProfile{
		ID:         model.UserId(userID.String()),
		Name:       info.Name.Value,
		AvatarPath: info.AvatarPath.Value,
		Bio:        info.Bio.Value,
		IsPublic:   info.IsPublic.Value,
	}
}

func checkEmailPattern(email string) bool {
	return strings.Contains(email, " ")
}
