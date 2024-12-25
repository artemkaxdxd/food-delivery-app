package userorder

import (
	userorder "backend/internal/entity/user_order"
	"time"
)

type UserOrder struct {
	Address     string `json:"address" valid:"required,stringlength(5|50)"`
	Comment     string `json:"comment"`
	PhoneNumber string `json:"phone_number" valid:"required"`
	Email       string `json:"email" valid:"required,email"`
}

func (u UserOrder) ToEntity(userID uint) userorder.UserOrder {
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

func (u UserOrder) ToEntityUpdate(orderID uint) userorder.UserOrder {
	return userorder.UserOrder{
		ID:          orderID,
		Address:     u.Address,
		Comment:     u.Comment,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
}
