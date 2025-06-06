package subscriptions

import (
	"user/internal/client/db"
	"user/internal/repository"
)

type repo struct {
	db db.Client
}

const (
	TableFollowsName = "follows"

	IdColumn          = "id"
	FollowerIdColumn  = "follower_id"  //последователь, подписчик
	FollowingIdColumn = "following_id" // тот на которого подписан
	StatusColumn      = "status"
	CreatedAtColumn   = "created_at"
	UpdatedAtColumn   = "updated_at"
)

const (
	tableUserName = "users"

	idUserColumn   = "id"
	userNameColumn = "username"
	avatarColumn   = "avatar_path"
)

func NewRepository(db db.Client) repository.SubscriptionsRepository {
	return &repo{db: db}
}
