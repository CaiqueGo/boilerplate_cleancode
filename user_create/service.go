package usercreate

import (
	"context"

	"github.com/boilerplate_cleancode/internal/log"
	"github.com/boilerplate_cleancode/internal/user"
)

type ServiceImpl struct {
	logger   log.Logger
	userRepo Repository
}

func NewService(logger log.Logger, userRepo Repository) *ServiceImpl {
	return &ServiceImpl{logger: logger, userRepo: userRepo}
}

func (r *ServiceImpl) Create(ctx context.Context, user user.User) (user.User, error) {
	r.logger.Info(ctx, "Payload User", "user", user)
	return r.userRepo.Create(ctx, user)
}
