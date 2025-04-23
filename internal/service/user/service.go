package user

import (
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

func (s serv) SaveUser(profile *model.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (s serv) UserProfile(id model.UserId) *model.Profile {
	//TODO implement me
	panic("implement me")
}

func (s serv) UpdateProfile(id model.UserId, updateInfo *model.UpdateProfile) *model.Profile {
	//TODO implement me
	panic("implement me")
}

func (s serv) Settings(id model.UserId) *model.Settings {
	//TODO implement me
	panic("implement me")
}
