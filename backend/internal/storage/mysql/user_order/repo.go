package userorder

import (
	userorder "backend/internal/entity/user_order"
	"backend/pkg/db"
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type repo struct {
	db db.Database
}

func NewRepo(db db.Database) repo {
	return repo{db}
}

func (r repo) NewTransaction(ctx context.Context) *gorm.DB {
	return r.db.Instance().WithContext(ctx).Begin()
}

func (r repo) GetOrderByID(ctx context.Context, orderID uint) (order userorder.UserOrder, err error) {
	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM user_orders WHERE id = $1`, orderID).Scan(&order).Error
	return
}

func (r repo) GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) (orders []userorder.UserOrder, err error) {
	rows, err := r.db.Instance().WithContext(ctx).Raw(`
		SELECT 
			o.id, o.user_id, o.address, o.comment, 
			o.phone_number, o.email, o.status, o.created_at,
			i.id, i.user_order_id, i.item_id, i.amount
		FROM 
			user_orders o
		LEFT JOIN 
			user_order_items i ON o.id = i.user_order_id
		WHERE 
			o.user_id = $1
		ORDER BY 
			o.id DESC
		LIMIT $2 OFFSET $3`, userID, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		orderMap     = make(map[uint]*userorder.UserOrder)
		orderItemIDs map[uint]struct{}
	)
	for rows.Next() {
		var (
			order = &userorder.UserOrder{}

			userOrderID, itemID, itemItemID, amount sql.NullInt64
		)

		err := rows.Scan(
			&order.ID, &order.UserID, &order.Address, &order.Comment,
			&order.PhoneNumber, &order.Email, &order.Status, &order.CreatedAt,
			&itemID, &userOrderID, &itemItemID, &amount,
		)
		if err != nil {
			return nil, err
		}

		if v, ok := orderMap[order.ID]; ok {
			order = v
		} else {
			orderMap[order.ID] = order
			orderItemIDs = make(map[uint]struct{})
		}

		if itemID.Valid {
			id := uint(itemID.Int64)
			if _, ok := orderItemIDs[id]; !ok {
				orderMap[order.ID].Items = append(orderMap[order.ID].Items, userorder.UserOrderItem{
					ID:          uint(itemID.Int64),
					UserOrderID: uint(userOrderID.Int64),
					ItemID:      uint(itemItemID.Int64),
					Amount:      uint(amount.Int64),
				})

				orderItemIDs[id] = struct{}{}
			}
		}
	}

	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders, nil
}

func (r repo) CreateOrder(ctx context.Context, tx *gorm.DB, body userorder.UserOrder) (userorder.UserOrder, error) {
	err := tx.WithContext(ctx).Create(&body).Error
	return body, err
}

func (r repo) CreateOrderItems(ctx context.Context, tx *gorm.DB, items []userorder.UserOrderItem) ([]userorder.UserOrderItem, error) {
	err := tx.WithContext(ctx).Create(&items).Error
	return items, err
}

func (r repo) UpdateOrder(ctx context.Context, body userorder.UserOrder) error {
	return r.db.Instance().WithContext(ctx).Where("id", body.ID).Updates(&body).Error
}

func (r repo) UpdateOrderStatus(ctx context.Context, orderID uint, status userorder.OrderStatus) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		UPDATE user_orders SET status = $1 WHERE id = $2`,
		status, orderID).Error
}

func (r repo) DeleteOrder(ctx context.Context, orderID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM user_orders WHERE id = $1`, orderID).Error
}

func (r repo) DeleteOrdersByUserID(ctx context.Context, userID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM user_orders WHERE user_id = $1`, userID).Error
}
