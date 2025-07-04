package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"time"
	"user/internal/client/db"
	"user/pkg/logger"
)

type pgClient struct {
	masterDBC db.DB
	dbc       *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	logger.Info("Connected to database", "dsn", dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	cl := &pgClient{
		masterDBC: NewDB(dbc),
		dbc:       dbc,
	}
	cl.StartPoolMonitoring(ctx, 10*time.Minute)

	return cl, nil
}

func (p *pgClient) DB() db.DB {
	return p.masterDBC
}

func (p *pgClient) Close() error {
	if p.masterDBC != nil {
		p.masterDBC.Close()
	}

	return nil
}
