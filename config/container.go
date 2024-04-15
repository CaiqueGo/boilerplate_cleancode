package config

import (
	"context"
	"database/sql"
)

type Config struct {
	Variable Environment
	Database *sql.DB
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

	return cfg, nil
}

func (r Config) Close() {
	r.Database.Close()
}
