package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (r *Config) DatabaseOpen() error {

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", r.Variable.dbUser, r.Variable.dbPass, r.Variable.dbHost, r.Variable.dbPort, r.Variable.dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	r.Database = db

	return nil
}
