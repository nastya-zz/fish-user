package user

import (
	"user/internal/client/db"
	"user/internal/client/minio/minio"
	"user/internal/repository"
	"user/internal/service"
)

type serv struct {
	userRepository  repository.UserRepository
	settingsService service.SettingsService
	txManager       db.TxManager
	minio           *minio.Client
}

func NewService(
	userRepository repository.UserRepository,
	settingsService service.SettingsService,
	txManager db.TxManager,
	minio *minio.Client,
) service.UserService {
	return &serv{
		userRepository:  userRepository,
		settingsService: settingsService,
		txManager:       txManager,
		minio:           minio,
	}
}
