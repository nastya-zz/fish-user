package user

import (
	"user/internal/client/db"
	"user/internal/repository"
)

const (
	tableName = "users"

	idColumn              = "id"
	nameColumn            = "username"
	emailColumn           = "email"
	avatarPathColumn      = "avatar_path"
	bioColumn             = "bio"
	isVerifiedColumn      = "is_verified"
	isPublicColumn        = "is_public"
	experienceLevelColumn = "experience_level"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

var column = map[string]string{
	"ID":              idColumn,
	"Name":            nameColumn,
	"Email":           emailColumn,
	"AvatarPath":      avatarPathColumn,
	"Bio":             bioColumn,
	"IsVerified":      isVerifiedColumn,
	"IsPublic":        isPublicColumn,
	"ExperienceLevel": experienceLevelColumn,
	"CreatedAt":       createdAtColumn,
	"UpdatedAt":       updatedAtColumn,
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}
