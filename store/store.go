package store

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"github.com/mizmorr/rest-example/store/migration"
	"github.com/mizmorr/rest-example/store/pg"
)

type Store struct {
	Pg *pg.DB
}

var (
	store Store
)

func New(ctx context.Context) (*Store, error) {

	pgDB, err := pg.Dial(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	//TODO: change to custom logger
	log.Println("Running pg migrations...")
	if err := migration.Run_migrations(); err != nil {
		return nil, errors.Wrap(err, "failed to run migrations")
	}
	if pgDB != nil {
		store.Pg = pgDB

	}
	return &store, nil
}
