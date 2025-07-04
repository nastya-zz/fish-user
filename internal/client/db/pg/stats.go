package pg

import (
	"context"
	"time"
	"user/pkg/logger"
)

// StartPoolMonitoring запускает мониторинг пула соединений
func (p *pgClient) StartPoolMonitoring(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				p.logPoolStats()
			}
		}
	}()
}

// logPoolStats логирует статистику пула соединений
func (p *pgClient) logPoolStats() {
	stats := p.dbc.Stat()
	logger.Info("PostgreSQL Pool Stats",
		"acquired_conns", stats.AcquiredConns(),
		"idle_conns", stats.IdleConns(),
		"max_conns", stats.MaxConns(),
		"total_conns", stats.TotalConns(),
		"new_conns_count", stats.NewConnsCount(),
		"acquire_count", stats.AcquireCount(),
		"acquire_duration", stats.AcquireDuration(),
		"canceled_acquire_count", stats.CanceledAcquireCount(),
	)
}
