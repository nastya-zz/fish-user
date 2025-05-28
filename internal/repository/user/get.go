package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"user/internal/client/db"
	"user/internal/logger"
	"user/internal/model"
)

func (r repo) UserProfile(ctx context.Context, id model.UserId) (*model.Profile, error) {
	const op = "user.UserProfile"

	uuId, err := model.GetUuid(id)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	builder := sq.Select(idColumn, nameColumn, emailColumn, isPublicColumn, createdAtColumn, experienceLevelColumn, isVerifiedColumn, avatarPathColumn).
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
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.IsPublic, &profile.CreatedAt, &profile.ExperienceLevel, &profile.IsVerified, &profile.AvatarPath)
	if err != nil {
		logger.Warn(op, fmt.Sprintf("error in get profile user %s", err.Error()))
		return nil, fmt.Errorf("error in get profile user %w", err)
	}
	logger.Info(op, "profile", profile)
	return &profile, nil
}
