package user

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"log"
	"user/internal/client/db"
	"user/internal/model"
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

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r repo) SaveUser(ctx context.Context, profile *model.Profile) error {
	const op = "user.SaveProfile"

	userID, err := uuid.Parse(string(profile.ID))

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, emailColumn, nameColumn, avatarPathColumn, bioColumn, isVerifiedColumn, createdAtColumn).
		Values(userID, profile.Email, profile.Name, "", "", profile.IsVerified, profile.CreatedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan()
	if err != nil {
		log.Println(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("error in save profile %w", err)
	}

	return nil
}

func (r repo) UserProfile(ctx context.Context, id model.UserId) *model.Profile {
	//TODO implement me
	panic("implement me")
}

func (r repo) UpdateProfile(ctx context.Context, id model.UserId, updateInfo *model.UpdateProfile) *model.Profile {
	//TODO implement me
	panic("implement me")
}
