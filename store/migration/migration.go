package migration

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/mizmorr/rest-example/config"
)

func Run_migrations() error {

	cfg := config.Get()
	if cfg.PgMigrationPath == "" {
		return fmt.Errorf("no migration path specified")
	}
	if cfg.PgURL == "" {
		return fmt.Errorf("no pg_url specified")
	}
	m, err := migrate.New(
		"file://"+cfg.PgMigrationPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
