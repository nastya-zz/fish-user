package user

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
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

func (r repo) SaveUser(ctx context.Context, profile *model.Profile) error {
	const op = "user.SaveProfile"

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, emailColumn, nameColumn, avatarPathColumn, bioColumn, isVerifiedColumn, createdAtColumn, updatedAtColumn).
		Values(profile.ID, profile.Email, profile.Name, "", "", profile.IsVerified, profile.CreatedAt, profile.CreatedAt)

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

func (r repo) UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error) {
	const op = "user.UserProfile"

	uuId, err := model.GetUuid(id)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	builder := sq.Select(idColumn, nameColumn, emailColumn, isPublicColumn, createdAtColumn, experienceLevelColumn, isVerifiedColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: uuId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var profile model.Profile
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.IsPublic, &profile.CreatedAt, &profile.ExperienceLevel, &profile.IsVerified)
	if err != nil {
		log.Println(fmt.Errorf("error in get profile user %w", err))
		return nil, fmt.Errorf("error in get profile user %w", err)
	}
	log.Println(op, profile)
	return &profile, nil
}

func (r repo) UpdateProfile(ctx context.Context, updateUser *model.UpdateProfile) (*model.Profile, error) {
	const op = "user.UpdateProfile"

	log.Printf("updating user %+v", updateUser)
	uuId, err := model.GetUuid(updateUser.ID)

	returning := fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s, %s ", idColumn, nameColumn, emailColumn, isPublicColumn, createdAtColumn, experienceLevelColumn, isVerifiedColumn)
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: uuId}).
		Suffix(returning)

	keys := [4]string{"Name", "AvatarPath", "Bio", "IsPublic"}
	fields := updateUser.GetFieldMap(*updateUser)

	for _, key := range keys {
		if val, ok := fields[key]; ok && val != nil {
			builder = builder.Set(column[key], fields[key])
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("cannot build query update user with id: %s", updateUser.ID)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}
	var profile model.Profile
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.IsPublic, &profile.CreatedAt, &profile.ExperienceLevel, &profile.IsVerified)
	if errors.Is(err, pgx.ErrNoRows) {
		log.Printf("error in update user with id: %s %+v", err, profile)

		return nil, fmt.Errorf("cannot update user with id: %s", updateUser.ID)
	}
	if err != nil {
		log.Printf("error in update user with id: %s", err)
		return nil, fmt.Errorf("cannot update user %w", err)
	}

	log.Printf("updating user %+v", updateUser)
	return &profile, nil
}
