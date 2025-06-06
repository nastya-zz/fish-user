package user

import (
	"context"
	"github.com/google/uuid"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/converter"
	"user/internal/model"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) GetProfile(ctx context.Context, req *desc.GetProfileRequest) (*desc.GetProfileResponse, error) {
	id := req.GetId()
	if id == "" || id == uuid.Nil.String() {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	uId, _ := uuid.Parse(id)

	profile, err := i.userService.UserProfile(ctx, model.UserId(uId.String()))
	if err != nil {
		return nil, status.Error(codes.NotFound, api_errors.UserIdNotFound)
	}

	return &desc.GetProfileResponse{Profile: converter.ToDescProfileFromProfile(profile)}, nil
}
