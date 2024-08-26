package migration

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
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
	_, err := migrate.New(
		cfg.PgMigrationPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	return nil
}
