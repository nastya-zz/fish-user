package settings

import (
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SettingsRepository {
	return &repo{db: db}
}

func (r repo) Settings(id model.UserId) *model.Settings {
	//TODO implement me
	panic("implement me")
}
