package user

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v4"
	"log"
	"user/internal/client/db"
	"user/internal/model"
)

func (r repo) SaveUser(ctx context.Context, profile *model.Profile) (model.UserId, error) {
	const op = "user.SaveProfile"

	uuId, err := model.GetUuid(profile.ID)
	if err != nil {
		return model.UserId(""), err
	}

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, emailColumn, nameColumn, avatarPathColumn, bioColumn, isVerifiedColumn, createdAtColumn, updatedAtColumn).
		Values(uuId, profile.Email, profile.Name, "", "", profile.IsVerified, profile.CreatedAt, profile.CreatedAt).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id model.UserId
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return id, nil
}
