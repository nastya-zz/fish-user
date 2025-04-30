package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
	"user/internal/client/db"
	"user/internal/model"
)

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
