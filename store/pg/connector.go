package pg

import (
	"time"

	"github.com/mizmorr/rest-example/config"
)

type Connector struct {
	Attempts int           `json:"connectAttempts"`
	Timeout  time.Duration `json:"connectTimeout"`
}

func Get_Connector() (c *Connector) {
	cfg := config.Get()
	return &Connector{cfg.PgConnAttempts, cfg.PgTimeout}
}
