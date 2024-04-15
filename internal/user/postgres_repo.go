package user

import (
	"context"
	"database/sql"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Create(ctx context.Context, user User) (User, error) {
	insertStatement := "INSERT INTO user (name) VALUES (?)"
	result, err := r.db.Exec(insertStatement, user.Name)

	if err != nil {
		return user, err
	}

	user.ID, err = result.LastInsertId()

	if err != nil {
		return user, err
	}

	return user, nil
}
