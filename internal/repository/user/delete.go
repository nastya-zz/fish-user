package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"user/internal/client/db"
	"user/internal/model"
	"user/pkg/logger"
)

func (r repo) DeleteUser(ctx context.Context, id model.UserId) error {
	const op = "user.Repository.Delete"

	builder := sq.Delete("users").Where(sq.Eq{"id": id}).Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var deletedId string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&deletedId)
	if err != nil {
		logger.Error(op, "error in delete user", "err", err)
		return fmt.Errorf("error in delete user %s,  %w", op, err)
	}

	return nil
}
