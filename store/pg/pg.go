package pg

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mizmorr/rest-example/config"
)

type DB struct {
	*pgxpool.Pool
}

var (
	once       sync.Once
	pgInstance *DB
)

func Dial(ctx context.Context) (*DB, error) {
	con := Get_Connector()
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, fmt.Errorf("no pg_url specified")
	}
	poolConfig, err := pgxpool.ParseConfig(cfg.PgURL)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}
	poolConfig.MaxConnIdleTime = cfg.PgMaxIdleTime
	poolConfig.HealthCheckPeriod = cfg.PgHealthCheckPeriod

	once.Do(func() {
		var db *pgxpool.Pool
		for con.Attempts > 0 {
			db, err = pgxpool.NewWithConfig(ctx, poolConfig)
			if err == nil {
				break
			}

			log.Printf("Postgres is trying to connect, attempts left: %d", con.Attempts)

			time.Sleep(con.Timeout)

			con.Attempts--
		}

		if err != nil {
			panic(fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err))
		}
		pgInstance = &DB{db}
	})
	if err := pgInstance.Ping(ctx); err != nil {
		return nil, err
	}
	return pgInstance, nil
}

func (db *DB) Close() {
	if db != nil {
		db.Close()
	}
}
