package pg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mizmorr/rest-example/config"
)

type DB struct {
	pg *pgxpool.Pool
}

var (
	once       sync.Once
	pgInstance *DB
)

func Dial(ctx context.Context) (*DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, fmt.Errorf("no pg_url specified")
	}
	once.Do(func() {
		db, err := pgxpool.New(ctx, cfg.PgURL)
		if err != nil {
			panic(err)

		}
		pgInstance = &DB{db}
	})
	if err := pgInstance.Ping(ctx); err != nil {
		return nil, err
	}
	return pgInstance, nil
}
func (db *DB) Ping(ctx context.Context) error {
	return db.pg.Ping(ctx)
}

func (db *DB) Close() {
	db.pg.Close()
}
