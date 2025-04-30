package user

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"log"
	"user/internal/client/db"
	"user/internal/model"
)

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
