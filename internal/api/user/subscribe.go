package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/model"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) Subscribe(ctx context.Context, req *desc.SubscribeRequest) (*emptypb.Empty, error) {
	userId := req.GetUserId()
	subscId := req.GetSubscriptionId()

	if subscId == "" || userId == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	if subscId == userId {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, api_errors.UserSubscribeCannotBeSame)
	}

	err := i.subscriptionsService.Subscribe(ctx, model.UserId(userId), model.UserId(subscId))
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
