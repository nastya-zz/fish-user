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
	minio           service.MinioService
	eventRepository repository.EventRepository
}

func NewService(
	userRepository repository.UserRepository,
	settingsService service.SettingsService,
	txManager db.TxManager,
	minio service.MinioService,
	eventRepository repository.EventRepository,
) service.UserService {
	return &serv{
		userRepository:  userRepository,
		settingsService: settingsService,
		txManager:       txManager,
		minio:           minio,
		eventRepository: eventRepository,
	}
}
