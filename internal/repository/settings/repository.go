package settings

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"user/internal/client/db"
	"user/pkg/logger"
	"user/internal/model"
	"user/internal/repository"
)

const (
	tableName          = "settings"
	idColumn           = "user_id"
	langColumn         = "language"
	availabilityColumn = "availability"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SettingsRepository {
	return &repo{db: db}
}

func (r repo) Create(ctx context.Context, id model.UserId, settings model.Settings) (model.UserId, error) {
	const op = "settings.Create"

	uuId, err := model.GetUuid(id)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, langColumn, availabilityColumn).
		Values(uuId, settings.Language, settings.Availability).
		Suffix("RETURNING user_id")

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var userId model.UserId
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userId)
	if err != nil {
		logger.Warn(op, "err", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return userId, nil
}

func (r repo) Get(ctx context.Context, id model.UserId) (*model.Settings, error) {
	const op = "settings.Get"

	uuId, err := model.GetUuid(id)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	builder := sq.Select(langColumn, availabilityColumn).
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

	var settings model.Settings

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&settings.Language, &settings.Availability)
	if err != nil {
		logger.Warn(fmt.Sprintf("error in get settings user %s", err.Error()))
		return nil, fmt.Errorf("error in get settings user %w", err)
	}

	logger.Info(op, settings)
	return &settings, nil
}

func (r repo) Update(ctx context.Context, id model.UserId, settings *model.Settings) (*model.Settings, error) {
	const op = "settings.Update"

	uuId, err := model.GetUuid(id)
	returning := fmt.Sprintf("RETURNING %s, %s ", langColumn, availabilityColumn)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(langColumn, settings.Language).
		Set(availabilityColumn, settings.Availability).
		Where(sq.Eq{idColumn: uuId}).
		Suffix(returning)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("cannot build query update settings with user id: %s", id)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}
	var updatedSettings model.Settings
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&updatedSettings.Language, &updatedSettings.Availability)

	if err != nil {
		logger.Warn("error in update settings with user id: %s", err)
		return nil, fmt.Errorf("cannot update settings %w", err)
	}

	return &updatedSettings, nil
}

func (r repo) Reset(ctx context.Context, id model.UserId) (*model.Settings, error) {
	const op = "repository.settings.Reset"

	logger.Info(op)

	defaultSettings := model.NewDefaultSettings()
	return r.Update(ctx, id, &defaultSettings)
}
