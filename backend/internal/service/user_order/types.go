package user

import (
	userorder "backend/internal/entity/user_order"
	"backend/pkg/logger"
	"context"
)

type (
	repo interface {
		GetOrderByID(ctx context.Context, orderID uint) (order userorder.UserOrder, err error)
		GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) (orders []userorder.UserOrder, err error)
		CreateOrder(ctx context.Context, body userorder.UserOrder) (userorder.UserOrder, error)
		UpdateOrder(ctx context.Context, body userorder.UserOrder) error
		DeleteOrder(ctx context.Context, orderID uint) error
	}

	service struct {
		repo repo
		l    logger.Logger
	}
)

func NewService(repo repo, l logger.Logger) service {
	return service{repo, l}
}
