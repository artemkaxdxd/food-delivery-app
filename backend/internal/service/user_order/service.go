package user

import (
	"backend/config"
	request "backend/internal/controllers/http/request/user_order"
	response "backend/internal/controllers/http/response/user_order"
	entity "backend/internal/entity/user_order"
	"context"
)

func (s service) GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) ([]response.Order, config.ServiceCode, error) {
	orders, err := s.repo.GetOrdersByUserID(ctx, userID, limit, offset)
	return response.OrdersToResponse(orders), config.DBErrToServiceCode(err), err
}

func (s service) CreateOrder(ctx context.Context, userID uint, body request.Order) (response.Order, config.ServiceCode, error) {
	if len(body.Items) == 0 {
		return response.Order{}, config.CodeEmptyOrder, config.ErrEmptyOrderItems
	}
	// TODO: add checks for duplicate item ids

	tx := s.repo.NewTransaction(ctx)
	defer tx.Rollback()

	order, err := s.repo.CreateOrder(ctx, tx, body.ToEntity(userID))
	if err != nil {
		return response.Order{}, config.DBErrToServiceCode(err), err
	}

	items := request.ItemsToEntity(body.Items, order.ID)
	if items, err = s.repo.CreateOrderItems(ctx, tx, items); err != nil {
		return response.Order{}, config.DBErrToServiceCode(err), err
	}

	err = tx.Commit().Error
	return response.OrderToResponse(order, items), config.DBErrToServiceCode(err), err
}

func (s service) UpdateOrder(ctx context.Context, userID, orderID uint, body request.Order) (config.ServiceCode, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return config.DBErrToServiceCode(err), err
	}
	if order.UserID != userID {
		return config.CodeForbidden, config.ErrInvalidOrderOwner
	}

	err = s.repo.UpdateOrder(ctx, body.ToEntityUpdate(orderID))
	return config.DBErrToServiceCode(err), err
}

func (s service) UpdateOrderStatus(ctx context.Context, userID, orderID uint, status entity.OrderStatus) (config.ServiceCode, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return config.DBErrToServiceCode(err), err
	}

	// Status should change the store admin, but for now let it be user
	if order.UserID != userID {
		return config.CodeForbidden, config.ErrInvalidOrderOwner
	}

	// TODO: Add checks for statuses order

	err = s.repo.UpdateOrderStatus(ctx, orderID, status)
	return config.DBErrToServiceCode(err), err
}

func (s service) DeleteOrder(ctx context.Context, userID, orderID uint) (config.ServiceCode, error) {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return config.DBErrToServiceCode(err), err
	}
	if order.UserID != userID {
		return config.CodeForbidden, config.ErrInvalidOrderOwner
	}

	err = s.repo.DeleteOrder(ctx, orderID)
	return config.DBErrToServiceCode(err), err
}
