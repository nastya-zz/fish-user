package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/converter"
	"user/internal/model"
)

func (i *Implementation) UpdateProfile(ctx context.Context, req *desc.UpdateProfileRequest) (*desc.UpdateProfileResponse, error) {
	info := req.GetInfo()

	if len(info.GetId()) == 0 {
		return nil, status.Error(codes.FailedPrecondition, "User ID empty")
	}

	profile, err := i.userService.UpdateProfile(ctx, mappingUpdateProfile(info))

	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot update profile")
	}

	return &desc.UpdateProfileResponse{Profile: converter.ToDescProfileFromProfile(*profile)}, nil
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
