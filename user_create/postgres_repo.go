package usercreate

import (
	"context"
	"database/sql"

	"github.com/boilerplate_cleancode/internal/user"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Create(ctx context.Context, user user.User) (user.User, error) {
	insertStatement := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	err := r.db.QueryRow(insertStatement, user.Name).Scan(&user.ID)
	if err != nil {
		println(err.Error())
		return user, err
	}

	return user, nil
}
