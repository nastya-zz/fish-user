package user

import (
	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r repo) SaveUser(profile *model.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (r repo) UserProfile(id model.UserId) *model.Profile {
	//TODO implement me
	panic("implement me")
}

func (r repo) UpdateProfile(id model.UserId, updateInfo *model.UpdateProfile) *model.Profile {
	//TODO implement me
	panic("implement me")
}
