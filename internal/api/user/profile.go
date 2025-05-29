package user

import (
	"context"
	"github.com/google/uuid"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/converter"
	"user/internal/model"
)

func (i *Implementation) GetProfile(ctx context.Context, req *desc.GetProfileRequest) (*desc.GetProfileResponse, error) {
	id := req.GetId()
	if id == "" || id == uuid.Nil.String() {
		return nil, status.Error(codes.InvalidArgument, "Не указан id пользователя")
	}

	uId, _ := uuid.Parse(id)

	profile, err := i.userService.UserProfile(ctx, model.UserId(uId.String()))
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь с таким Id не найден")
	}

	return &desc.GetProfileResponse{Profile: converter.ToDescProfileFromProfile(profile)}, nil
}
