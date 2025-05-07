package subscriptions

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
)

func (r repo) Unsubscribe(ctx context.Context, id model.UserId, subscriptionId model.UserId) error {
	const op = "subscriptions.Unsubscribe"

	userId, err := model.GetUuid(id)
	subsId, err := model.GetUuid(subscriptionId)
	if err != nil {
		return err
	}

	builder := sq.Delete(TableFollowsName).
		Where(sq.Eq{FollowerIdColumn: userId, FollowingIdColumn: subsId}).
		Limit(1).
		Suffix("RETURNING id")

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
		log.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	log.Printf(" rowId %s, subscriptionId %s,  id %s", rowId, subscriptionId, id)

	return nil
}
