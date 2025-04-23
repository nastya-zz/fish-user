package model

import "time"

type UserId string

type Profile struct {
	ID              UserId
	Name            string
	Email           string
	AvatarPath      string
	Bio             string
	ExperienceLevel int
	IsVerified      bool
	IsPublic        bool
	CreatedAt       time.Time
}

type UpdateProfile struct {
	ID         UserId
	Name       string
	Email      string
	AvatarPath string
	Bio        string
	IsPublic   bool
}
