package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/model"
)

func (i *Implementation) UnSubscribe(ctx context.Context, req *desc.SubscribeRequest) (*emptypb.Empty, error) {
	userId := req.GetUserId()
	subscId := req.GetSubscriptionId()
	if subscId == "" || userId == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "invalid arguments")
	}

	err := i.subscriptionsService.Unsubscribe(ctx, model.UserId(userId), model.UserId(subscId))
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
