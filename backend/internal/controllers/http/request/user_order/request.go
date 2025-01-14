package userorder

import (
	userorder "backend/internal/entity/user_order"
	"time"
)

type (
	Order struct {
		Address     string      `json:"address" valid:"required,stringlength(5|50)"`
		Comment     string      `json:"comment"`
		PhoneNumber string      `json:"phone_number" valid:"required"`
		Email       string      `json:"email" valid:"required,email"`
		Items       []OrderItem `json:"items"`
	}

	OrderItem struct {
		ItemID uint `json:"item_id"`
		Amount uint `json:"amount"`
	}

	OrderStatus struct {
		Status userorder.OrderStatus `json:"status" valid:"required"`
	}
)

func (u Order) ToEntity(userID uint) userorder.UserOrder {
	return userorder.UserOrder{
		UserID:      userID,
		Address:     u.Address,
		Comment:     u.Comment,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Status:      userorder.OrderStatusCreated,
		CreatedAt:   time.Now().Unix(),
	}
}

func (u Order) ToEntityUpdate(orderID uint) userorder.UserOrder {
	return userorder.UserOrder{
		ID:          orderID,
		Address:     u.Address,
		Comment:     u.Comment,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
}

func (u OrderItem) ToEntity(orderID uint) userorder.UserOrderItem {
	return userorder.UserOrderItem{
		UserOrderID: orderID,
		ItemID:      u.ItemID,
		Amount:      u.Amount,
	}
}

func ItemsToEntity(items []OrderItem, orderID uint) []userorder.UserOrderItem {
	if len(items) == 0 {
		return nil
	}

	resp := make([]userorder.UserOrderItem, len(items))
	for i, v := range items {
		resp[i] = v.ToEntity(orderID)
	}
	return resp
}
