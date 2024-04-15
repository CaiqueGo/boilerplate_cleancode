package usercreate

import (
	"context"

	"github.com/boilerplate_cleancode/internal/user"
)

type ServiceImpl struct {
	userRepo Repository
}

func NewService(userRepo Repository) *ServiceImpl {
	return &ServiceImpl{userRepo: userRepo}
}

func (r *ServiceImpl) Create(ctx context.Context, user user.User) (user.User, error) {
	return r.userRepo.Create(ctx, user)
}
