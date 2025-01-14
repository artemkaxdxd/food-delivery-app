package user

import (
	userorder "backend/internal/entity/user_order"
	"backend/pkg/logger"
	"context"

	"gorm.io/gorm"
)

type (
	repo interface {
		NewTransaction(ctx context.Context) *gorm.DB

		GetOrderByID(ctx context.Context, orderID uint) (order userorder.UserOrder, err error)
		GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) (orders []userorder.UserOrder, err error)

		CreateOrder(ctx context.Context, tx *gorm.DB, body userorder.UserOrder) (userorder.UserOrder, error)
		CreateOrderItems(ctx context.Context, tx *gorm.DB, items []userorder.UserOrderItem) ([]userorder.UserOrderItem, error)

		UpdateOrder(ctx context.Context, body userorder.UserOrder) error
		UpdateOrderStatus(ctx context.Context, orderID uint, status userorder.OrderStatus) error

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
