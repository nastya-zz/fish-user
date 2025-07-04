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

	uuid, err := model.GetUuid(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	builder := sq.Delete("users").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": uuid})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		logger.Error(op, "error in delete user", "err", err)
		return fmt.Errorf("error in delete user %s,  %w", op, err)
	}

	return nil
}
