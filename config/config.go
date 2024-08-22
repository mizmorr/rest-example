package config

import (
	"sync"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	LogLevel        string `yaml:"log_level"`
	PgURL           string `yaml:"pg_URL"`
	PgMigrationPath string `yaml:"pg_migration_path"`
	HTTPAddress     string `yaml:"http_address"`
	FilePath        string `yaml:"file_path"`
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
		err := config.LoadFiles("config.yml")
		if err != nil {
			panic(err)
		}
		err = config.BindStruct("config", conf)
		if err != nil {
			panic(err)
		}
	})
	return &conf

}
