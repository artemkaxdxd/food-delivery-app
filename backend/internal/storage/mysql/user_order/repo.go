package userorder

import (
	userorder "backend/internal/entity/user_order"
	"backend/pkg/db"
	"context"
)

type repo struct {
	db db.Database
}

func NewRepo(db db.Database) repo {
	return repo{db}
}

func (r repo) GetOrderByID(ctx context.Context, orderID uint) (order userorder.UserOrder, err error) {
	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM orders WHERE id = $1`, orderID).Scan(&order).Error
	return
}

func (r repo) GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) (orders []userorder.UserOrder, err error) {
	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM orders WHERE user_id = $1`, userID).Scan(&orders).Error
	return
}

func (r repo) CreateOrder(ctx context.Context, body userorder.UserOrder) (userorder.UserOrder, error) {
	err := r.db.Instance().WithContext(ctx).Create(&body).Error
	return body, err
}

func (r repo) UpdateOrder(ctx context.Context, body userorder.UserOrder) error {
	return r.db.Instance().WithContext(ctx).Where("id", body.ID).Updates(&body).Error
}

func (r repo) DeleteOrder(ctx context.Context, orderID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM user_orders WHERE id = $1`, orderID).Error
}

func (r repo) DeleteOrdersByUserID(ctx context.Context, userID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM user_orders WHERE user_id = $1`, userID).Error
}
