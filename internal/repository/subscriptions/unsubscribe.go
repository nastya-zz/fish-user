package subscriptions

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"user/internal/client/db"
	"user/internal/logger"
	"user/internal/model"
)

func (r repo) Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	const op = "subscriptions.Unsubscribe"

	userId, err := model.GetUuid(id)
	subsId, err := model.GetUuid(subscriptionId)
	if err != nil {
		return err
	}

	builder := sq.Delete(TableFollowsName).
		Where(sq.And{
			sq.Eq{FollowerIdColumn: subsId},
			sq.Eq{FollowingIdColumn: userId},
		}).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("cannot build query update user with id: %s", userId)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var rowId uuid.UUID
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&rowId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Error(op, "err", fmt.Errorf("В данной связке пользователей подписка не найдена userId: %s, subsId: %s", userId, subsId))
			return fmt.Errorf("В данной связке пользователей подписка не найдена userId: %s, subsId: %s", userId, subsId)
		}

		logger.Error(op, "err", err)
		return err
	}

	logger.Warn(op, "rowId", rowId, "subscriptionId", subscriptionId, "id", id)

	return nil
}
