package service

import (
	"context"
	"user/internal/model"
)

type UserService interface {
	SaveUser(context.Context, *model.Profile) (model.UserId, error)
	UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error)
	UpdateProfile(ctx context.Context, updateInfo *model.UpdateProfile) (*model.Profile, error)
	UpdateInfo(ctx context.Context, updateInfo *model.UpdateUser) error
	UploadAvatar(ctx context.Context, file []byte, name string, userId model.UserId) (string, error)
	RemoveAvatar(ctx context.Context, userId model.UserId, filename string) error
	//BlockUser(ctx context.Context, id model.UserId) error
	DeleteUser(ctx context.Context, id model.UserId) error
}

type SettingsService interface {
	Create(ctx context.Context, id model.UserId) (model.UserId, error)
	Get(ctx context.Context, id model.UserId) (*model.Settings, error)
	Update(ctx context.Context, id model.UserId, settings *model.Settings) (*model.Settings, error)
	Reset(ctx context.Context, id model.UserId) (*model.Settings, error)
}

type SubscriptionsService interface {
	Subscriptions(ctx context.Context, id model.UserId) (*model.Subscriptions, error)
	Subscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error
	Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error
}

type EventsService interface {
	Process(ctx context.Context, event model.Event) error
}

type MinioService interface {
	UploadFile(ctx context.Context, file []byte, filename string) (string, error)
	RemoveFile(ctx context.Context, link string) error
}
