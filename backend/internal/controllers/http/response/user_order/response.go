package userorder

import (
	userorder "backend/internal/entity/user_order"
)

type UserOrder struct {
	ID          uint                  `json:"id"`
	UserID      uint                  `json:"user_id"`
	Address     string                `json:"address"`
	Comment     string                `json:"comment"`
	PhoneNumber string                `json:"phone_number"`
	Email       string                `json:"email"`
	Status      userorder.OrderStatus `json:"status"`
	CreatedAt   int64                 `json:"created_at"`
}

func OrderToResponse(order userorder.UserOrder) UserOrder {
	return UserOrder(order)
}

func OrdersToResponse(orders []userorder.UserOrder) []UserOrder {
	resp := make([]UserOrder, len(orders))
	for i, v := range orders {
		resp[i] = OrderToResponse(v)
	}
	return resp
}
