package user

import (
	"context"

	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"user/internal/model"
	"user/internal/utils"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) UnSubscribe(ctx context.Context, req *desc.SubscribeRequest) (*emptypb.Empty, error) {
	userID, err := utils.GetUserIdFromMetadata(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	subscId := req.GetSubscriptionId()
	if subscId == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	err = i.subscriptionsService.Unsubscribe(ctx, model.UserId(userID.String()), model.UserId(subscId))
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, api_errors.UserUnsubscribeFailed)
	}

	return &emptypb.Empty{}, nil
}
