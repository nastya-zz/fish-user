package settings

import (
	"context"
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

func (s serv) Create(ctx context.Context, id model.UserId) (model.UserId, error) {
	defaultSettings := model.NewDefaultSettings()
	return s.settingsRepository.Create(ctx, id, defaultSettings)
}

func (s serv) Get(ctx context.Context, id model.UserId) (*model.Settings, error) {
	return s.settingsRepository.Get(ctx, id)
}

func (s serv) Update(ctx context.Context, id model.UserId, settings *model.Settings) (*model.Settings, error) {
	return s.settingsRepository.Update(ctx, id, settings)
}

func (s serv) Reset(ctx context.Context, id model.UserId) (*model.Settings, error) {
	return s.settingsRepository.Reset(ctx, id)
}
