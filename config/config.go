package config

import (
	"sync"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	LogLevel            string        `yaml:"log_level"`
	PgURL               string        `yaml:"pg_URL"`
	PgMigrationPath     string        `yaml:"pg_migration_path"`
	PgTimeout           time.Duration `yaml:"pg_timeout"`
	PgConnAttempts      int           `yaml:"pg_conn_attempts"`
	PgHealthCheckPeriod time.Duration `yaml:"pg_health_check_period"`
	PgMaxIdleTime       time.Duration `yaml:"pg_max_idle_time"`
	HTTPAddress         string        `yaml:"http_address"`
	FilePath            string        `yaml:"file_path"`
}

var (
	once sync.Once
	conf Config
)

func Get() *Config {
	once.Do(func() {
		config.AddDriver(yaml.Driver)
		config.WithOptions(func(opt *config.Options) {
			opt.DecoderConfig.TagName = "yaml"
		})
		err := config.LoadFiles("../../config/config.yaml")
		if err != nil {
			panic(err)
		}
		err = config.BindStruct("", &conf)
		if err != nil {
			panic(err)
		}

	})
	return &conf

}
