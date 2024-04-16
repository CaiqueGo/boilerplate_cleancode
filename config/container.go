package config

import (
	"context"
	"database/sql"

	"github.com/boilerplate_cleancode/internal/log"
	lognative "github.com/boilerplate_cleancode/internal/log/log_native"
)

type Config struct {
	Variable Environment
	Database *sql.DB
	Log      log.Logger
}

func Start(ctx context.Context) (*Config, error) {
	env, err := Load()
	if err != nil {
		return &Config{}, err
	}

	cfg := &Config{Variable: env}
	err = cfg.DatabaseOpen()
	if err != nil {
		return cfg, err
	}

	cfg.Log = lognative.NewLogNative()

	return cfg, nil
}

func (r Config) Close() {
	r.Database.Close()
}
