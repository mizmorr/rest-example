package store

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"

	repo "github.com/mizmorr/rest-example/internal/repository"
	"github.com/mizmorr/rest-example/pkg/logger"
	"github.com/mizmorr/rest-example/store/migration"
	"github.com/mizmorr/rest-example/store/pg"
)

type Store struct {
	Pg   *pg.DB
	User UserRepo
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
		go store.KeepAlive(ctx)
		store.User = repo.NewUserRepo(pgDB)
	}
	return &store, nil
}

const KeepALiveTimeout = 5

func (store *Store) KeepAlive(ctx context.Context) {
	logger := logger.Get()
	for {
		time.Sleep(time.Second * KeepALiveTimeout)
		var (
			lost_connection bool
			err             error
		)

		if store.Pg == nil {
			lost_connection = true
		} else if err := store.Pg.Ping(ctx); err != nil {
			lost_connection = true
		}
		if lost_connection {
			logger.Debug().Msg("[store.KeepAlivePg] Lost connection, is trying to reconnect...")
			store.Pg, err = pg.Dial(ctx)
			if err != nil {
				logger.Err(err)
			} else {
				logger.Debug().Msg("[store.KeepAlivePg] Connection established")
			}
		}

	}
}
