package user

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"user/internal/model"
)

func (i *Implementation) GetSubscriptions(ctx context.Context, req *desc.GetSubscriptionsRequest) (*desc.GetSubscriptionsResponse, error) {
	id := req.GetId()
	if len(id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing user id")
	}

	subs, err := i.subscriptionsService.Subscriptions(ctx, model.UserId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Println("subscriptions:", subs)

	return &desc.GetSubscriptionsResponse{
		Subscriptions: mapping(subs.Subscriptions),
		Subscribers:   mapping(subs.Subscribers),
	}, nil

}

func mapping(sb []model.Subscription) []*desc.SubscriptionUser {
	log.Printf("mapping %d", len(sb))

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
	log.Printf("slice %d", len(slice))

	return slice
}
