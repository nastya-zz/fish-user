package model

import "time"

type Event struct {
	ID      int    `db:"id"`
	Type    string `db:"event_type"`
	Payload []byte `db:"payload"`
}

const (
	UserCreate        = "user_create"
	UserUpdate        = "user_update"
	UserUpdateProfile = "user_update_profile"
	UserDelete        = "user_delete"
)

type UserPayload struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"isVerified"`
	CreatedAt  time.Time `json:"createdAt"`
}
