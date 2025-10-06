package event

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"user/internal/client/db"
	"user/internal/model"
	"user/internal/repository"
	"user/pkg/logger"
)

const (
	tableName = "events"

	idColumn        = "id"
	eventTypeColumn = "event_type"
	payloadColumn   = "payload"
	statusColumn    = "status"
	createdAtColumn = "created_at"
)

const (
	newStatus  = "new"
	doneStatus = "done"
)

type repo struct {
	db db.Client
}

type event struct {
	ID      int    `db:"id"`
	Type    string `db:"event_type"`
	Payload string `db:"payload"`
}

func NewRepository(db db.Client) repository.EventRepository {
	return &repo{db: db}
}

func (r repo) GetNewEvent(ctx context.Context, batchSize int) ([]*model.Event, error) {
	const op = "db.GetNewEvent"

	builder := sq.Select(idColumn, eventTypeColumn, payloadColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{statusColumn: newStatus}).
		Limit(uint64(batchSize))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var events []*model.Event
	err = r.db.DB().ScanAllContext(ctx, &events, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}

func (r repo) SaveEvent(ctx context.Context, event *model.Event) error {
	const op = "db.SaveEvent"

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(eventTypeColumn, payloadColumn).
		Values(event.Type, event.Payload).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var savedId int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&savedId)
	if err != nil {
		logger.Warn(op, "error", fmt.Errorf("error in create event %w", err))
		return fmt.Errorf("error in create event %w", err)
	}

	return nil
}

func (r repo) SetDone(ctx context.Context, id int) error {
	const op = "db.SetDone"
	logger.Info(op, "updating event ", id)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(statusColumn, doneStatus).
		Where(sq.Eq{idColumn: id}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("cannot build query event with id: %d", id)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var savedId int
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&savedId)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("error in update event with id: ", "err", err, "savedId", savedId)

		return fmt.Errorf("cannot update event with id: %d", savedId)
	}
	if err != nil {
		logger.Warn("error in update event with id: ", "err", err)
		return fmt.Errorf("cannot update event %w", err)
	}

	return nil
}
