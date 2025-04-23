package model

type Subscription struct {
	ID         UserId
	Name       string
	AvatarPath string
}

type Subscriptions struct {
	Followers   []Subscription
	Subscribers []Subscription
}
