package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/logger"
	"user/internal/model"
	api_errors "user/pkg/api-errors"
)

func (i *Implementation) GetSubscriptions(ctx context.Context, req *desc.GetSubscriptionsRequest) (*desc.GetSubscriptionsResponse, error) {
	const op = "api.GetSubscriptions"
	id := req.GetId()
	if len(id) == 0 {
		return nil, status.Error(codes.InvalidArgument, api_errors.UserIdRequired)
	}

	subs, err := i.subscriptionsService.Subscriptions(ctx, model.UserId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.Debug(op, "subscriptions:", subs)

	return &desc.GetSubscriptionsResponse{
		Subscriptions: mapping(subs.Subscriptions),
		Subscribers:   mapping(subs.Subscribers),
	}, nil
}

func mapping(sb []model.Subscription) []*desc.SubscriptionUser {
	if len(sb) == 0 {
		return []*desc.SubscriptionUser{}
	}

	slice := make([]*desc.SubscriptionUser, 0, len(sb))

	for _, element := range sb {
		slice = append(slice, &desc.SubscriptionUser{
			Id:         string(element.ID),
			Name:       element.Name,
			AvatarPath: element.AvatarPath,
		})
	}

	return slice
}
