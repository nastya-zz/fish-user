package model

import (
	"github.com/google/uuid"
	"reflect"
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

type UpdateUser struct {
	ID         UserId
	Name       string
	Email      string
	IsVerified bool
}

// todo refactor исправить на динамический маппинг
func (up *UpdateProfile) GetFieldMap(u UpdateProfile) map[string]interface{} {
	return map[string]interface{}{
		"Name":       u.Name,
		"AvatarPath": u.AvatarPath,
		"Bio":        u.Bio,
		"IsPublic":   u.IsPublic,
	}
}

func (up *UpdateUser) GetFieldMap(u *UpdateUser, keys []string) map[string]interface{} {
	fieldsMap := make(map[string]interface{})

	v := reflect.ValueOf(u).Elem()

	for _, key := range keys {
		field := v.FieldByName(key)
		if field.IsValid() {
			fieldsMap[key] = field.Interface()
		}
	}
	return fieldsMap
}

func GetUuid[T ~string](id T) (uuid.UUID, error) {
	return uuid.Parse(string(id))
}
