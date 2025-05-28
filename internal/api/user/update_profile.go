package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"user/internal/converter"
	"user/internal/logger"
	"user/internal/model"
)

func (i *Implementation) UpdateProfile(ctx context.Context, req *desc.UpdateProfileRequest) (*desc.UpdateProfileResponse, error) {
	const op = "api.user.UpdateProfile"
	info := req.GetInfo()

	if len(info.GetId()) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "User ID empty")
	}

	if len(info.Email.Value) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "User email empty")
	} else if checkEmailPattern(info.Email.Value) {
		return nil, status.Error(codes.FailedPrecondition, "Email not match pattern")
	}

	profile, err := i.userService.UpdateProfile(ctx, mappingUpdateProfile(info))

	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot update profile")
	}

	logger.Info(op, "updating user ", profile)

	return &desc.UpdateProfileResponse{Profile: converter.ToDescProfileFromProfile(profile)}, nil
}

func mappingUpdateProfile(info *desc.UpdateProfile) *model.UpdateProfile {
	return &model.UpdateProfile{
		ID:         model.UserId(info.GetId()),
		Name:       info.Name.Value,
		AvatarPath: info.AvatarPath.Value,
		Bio:        info.Bio.Value,
		IsPublic:   info.IsPublic.Value,
	}
}

func checkEmailPattern(email string) bool {
	return strings.Contains(email, " ")
}
