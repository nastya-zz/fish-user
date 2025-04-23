package settings

import (
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/service"
)

type serv struct {
	settingsRepository repository.SettingsRepository
	txManager          db.TxManager
}

func NewService(
	settingsRepository repository.SettingsRepository,
	txManager db.TxManager,
) service.SettingsService {
	return &serv{
		settingsRepository: settingsRepository,
		txManager:          txManager,
	}
}

func (s serv) Settings(id model.UserId) *model.Settings {
	//TODO implement me
	panic("implement me")
}
