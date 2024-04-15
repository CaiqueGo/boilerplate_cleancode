package usercreate

import (
	"context"

	"github.com/boilerplate_cleancode/internal/user"
)

type Repository interface {
	Create(ctx context.Context, user user.User) (user.User, error)
}

type Service interface {
	Create(ctx context.Context, user user.User) (user.User, error)
}
