package subscriptions

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"user/internal/client/db"
	"user/pkg/logger"
	"user/internal/model"
)

func (r repo) Subscriptions(ctx context.Context, id model.UserId) (*model.Subscriptions, error) {
	const op = "user.Subscriptions"

	uuId, err := model.GetUuid(id)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	subscribers, err := r.getFollowers(ctx, uuId)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	subscriptions, err := r.getFollowing(ctx, uuId)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	logger.Info(op, "subscribers: ", subscribers)
	logger.Info(op, "subscriptions: ", subscriptions)

	return &model.Subscriptions{
		Subscribers:   subscribers,
		Subscriptions: subscriptions,
	}, nil

}

func (r repo) SubscriptionExists(ctx context.Context, id model.UserId, subscriptionId model.UserId) (bool, error) {
	const op = "user.SubscriptionExists"

	uuId, err := model.GetUuid(id)
	if err != nil {
		return false, fmt.Errorf("%s, %w", op, err)
	}
	subscUuId, err := model.GetUuid(subscriptionId)
	if err != nil {
		return false, fmt.Errorf("%s, %w", op, err)
	}

	builder := sq.Select("u.id, u.username, u.avatar_path").
		From("follows f").
		Join("users u ON f.follower_id = u.id").
		Where(sq.Eq{
			"f.following_id": uuId,
			"f.status":       "accepted",
			"f.follower_id":  subscUuId,
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var sbc model.Subscription
	if err = r.db.DB().ScanOneContext(ctx, &sbc, q, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		logger.Warn(fmt.Errorf("error in get sbsc user %w", err).Error())
		return false, fmt.Errorf(op+" error in get sbsc user %w", err)
	}

	return true, nil
}

func (r repo) getFollowers(ctx context.Context, id uuid.UUID) ([]model.Subscription, error) {
	const op = "user.getFollowers"
	/**
	  SELECT u.* FROM follows f
	  JOIN users u ON f.follower_id = u.id
	  WHERE f.following_id = :user_id AND f.status = 'accepted';
	*/

	builder := sq.Select("u.id, u.username, u.avatar_path").
		From("follows f").
		Join("users u ON f.follower_id = u.id").
		Where(sq.Eq{
			"f.following_id": id,
			"f.status":       "accepted",
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var sbc []model.Subscription

	if err = r.db.DB().ScanAllContext(ctx, &sbc, q, args...); err != nil {
		logger.Warn(fmt.Sprintf("error in get sbsc user %w", err))
		return nil, fmt.Errorf("error in get sbsc user %w", err)
	}

	logger.Info(op, "sbc", sbc)
	return sbc, nil

}
func (r repo) getFollowing(ctx context.Context, id uuid.UUID) ([]model.Subscription, error) {
	const op = "user.getFollowing"

	/**
	SELECT u.* FROM follows f
	JOIN users u ON f.following_id = u.id
	WHERE f.follower_id = :user_id AND f.status = 'accepted';
	*/

	builder := sq.Select("u.id, u.username, u.avatar_path").
		From("follows f").
		Join("users u ON f.following_id = u.id").
		Where(sq.Eq{
			"f.follower_id": id,
			"f.status":      "accepted",
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var sbc []model.Subscription

	if err = r.db.DB().ScanAllContext(ctx, &sbc, q, args...); err != nil {
		logger.Warn(fmt.Sprintf("error in get sbsc user %s", err.Error()))
		return nil, fmt.Errorf("error in get sbsc user %w", err)
	}

	logger.Info(op, "sbc", sbc)
	return sbc, nil

}
