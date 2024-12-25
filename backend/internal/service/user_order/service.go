package user

import (
	"backend/config"
	request "backend/internal/controllers/http/request/user_order"
	response "backend/internal/controllers/http/response/user_order"
	"context"
)

func (s service) GetOrdersByUserID(ctx context.Context, userID, limit, offset uint) ([]response.UserOrder, config.ServiceCode, error) {
	orders, err := s.repo.GetOrdersByUserID(ctx, userID, limit, offset)
	return response.OrdersToResponse(orders), config.DBErrToServiceCode(err), err
}

func (s service) CreateOrder(ctx context.Context, userID uint, body request.UserOrder) (response.UserOrder, config.ServiceCode, error) {
	order, err := s.repo.CreateOrder(ctx, body.ToEntity(userID))
	return response.OrderToResponse(order), config.DBErrToServiceCode(err), err
}

func (s service) UpdateOrder(ctx context.Context, userID, orderID uint, body request.UserOrder) (config.ServiceCode, error) {
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
