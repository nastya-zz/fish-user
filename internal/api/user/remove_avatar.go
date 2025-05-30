package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/logger"
	"user/internal/model"
)

func (i *Implementation) RemoveAvatar(ctx context.Context, req *desc.RemoveAvatarRequest) (*emptypb.Empty, error) {
	const op = "api.user.RemoveAvatar"

	userId := req.GetUserId()
	if len(userId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	filename := req.GetFilename()
	if len(filename) == 0 {
		return nil, status.Error(codes.InvalidArgument, "filename is required")
	}

	err := i.userService.RemoveAvatar(ctx, model.UserId(userId), filename)
	if err != nil {
		logger.Warn(op, "err", err.Error())
		return nil, status.Error(codes.Internal, "cannot remove avatar")
	}

	return &emptypb.Empty{}, nil
}
