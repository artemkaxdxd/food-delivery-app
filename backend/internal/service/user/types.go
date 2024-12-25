package user

import (
	"backend/config"
	"backend/internal/entity/user"
	"backend/pkg/logger"
	"context"
)

type (
	repo interface {
		GetUserByID(ctx context.Context, userID uint) (user.User, error)
		GetUserByEmail(ctx context.Context, email string) (user.User, error)

		CreateUser(ctx context.Context, body user.User) (user.User, error)
		UpdateUser(ctx context.Context, body user.User) error
		DeleteUser(ctx context.Context, userID uint) error
	}

	orderRepo interface {
		DeleteOrdersByUserID(ctx context.Context, userID uint) error
	}

	service struct {
		repo      repo
		orderRepo orderRepo
		l         logger.Logger
		cfgJWT    config.JWT
	}
)

func NewService(
	repo repo,
	orderRepo orderRepo,
	l logger.Logger,
	cfgJWT config.JWT,
) service {
	return service{repo, orderRepo, l, cfgJWT}
}
