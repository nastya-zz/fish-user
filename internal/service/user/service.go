package user

import (
	"user/internal/client/db"
	"user/internal/repository"
	"user/internal/service"
)

type serv struct {
	userRepository  repository.UserRepository
	settingsService service.SettingsService
	txManager       db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	settingsService service.SettingsService,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository:  userRepository,
		settingsService: settingsService,
		txManager:       txManager,
	}
}
