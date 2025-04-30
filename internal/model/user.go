package model

import (
	"github.com/google/uuid"
	"time"
)

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
	AvatarPath string
	Bio        string
	IsPublic   bool
}

func (up *UpdateProfile) GetFieldMap(u UpdateProfile) map[string]interface{} {
	return map[string]interface{}{
		"Name":       u.Name,
		"AvatarPath": u.AvatarPath,
		"Bio":        u.Bio,
		"IsPublic":   u.IsPublic,
	}
}

func GetUuid[T ~string](id T) (uuid.UUID, error) {
	return uuid.Parse(string(id))
}
