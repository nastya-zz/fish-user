package model

type Subscription struct {
	ID         UserId `db:"id"`
	Name       string `db:"username"`
	AvatarPath string `db:"avatar_path"`
}

type Subscriptions struct {
	Subscribers   []Subscription
	Subscriptions []Subscription
}
