package userorder

import (
	userorder "backend/internal/entity/user_order"
)

type (
	Order struct {
		ID          uint                  `json:"id"`
		UserID      uint                  `json:"user_id"`
		Address     string                `json:"address"`
		Comment     string                `json:"comment"`
		PhoneNumber string                `json:"phone_number"`
		Email       string                `json:"email"`
		Status      userorder.OrderStatus `json:"status"`
		CreatedAt   int64                 `json:"created_at"`

		Items []OrderItem `json:"order_items"`
	}

	OrderItem struct {
		ID          uint `json:"id"`
		UserOrderID uint `json:"user_order_id"`
		ItemID      uint `json:"item_id"`
		Amount      uint `json:"amount"`
	}
)

func OrderToResponse(order userorder.UserOrder, items []userorder.UserOrderItem) Order {
	itemsResp := ItemsToResponse(items)
	return Order{
		ID:          order.ID,
		UserID:      order.UserID,
		Address:     order.Address,
		Comment:     order.Comment,
		PhoneNumber: order.PhoneNumber,
		Email:       order.Email,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
		Items:       itemsResp,
	}
}

func OrdersToResponse(orders []userorder.UserOrder) []Order {
	if len(orders) == 0 {
		return nil
	}

	resp := make([]Order, len(orders))
	for i, v := range orders {
		resp[i] = OrderToResponse(v, v.Items)
	}
	return resp
}

func ItemToResponse(item userorder.UserOrderItem) OrderItem {
	return OrderItem{
		ID:          item.ID,
		UserOrderID: item.UserOrderID,
		ItemID:      item.ItemID,
		Amount:      item.Amount,
	}
}

func ItemsToResponse(items []userorder.UserOrderItem) []OrderItem {
	if len(items) == 0 {
		return nil
	}

	resp := make([]OrderItem, len(items))
	for i, v := range items {
		resp[i] = ItemToResponse(v)
	}
	return resp
}
