package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/model"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) UploadAvatar(ctx context.Context, req *desc.UploadAvatarRequest) (*desc.UploadAvatarResponse, error) {
	fileBytes := req.GetImage()
	if len(fileBytes) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserAvatarFileEmpty)
	}

	filename := req.GetFilename()
	if len(filename) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserAvatarFilenameEmpty)
	}

	userId := req.GetUserId()
	if len(userId) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	res, err := i.userService.UploadAvatar(ctx, fileBytes, filename, model.UserId(userId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UploadAvatarResponse{
		Link: res,
	}, nil
}
