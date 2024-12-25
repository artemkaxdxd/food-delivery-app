package userorder

type UserOrder struct {
	ID          uint
	UserID      uint
	Address     string
	Comment     string
	PhoneNumber string
	Email       string
	Status      OrderStatus
	CreatedAt   int64
}

func (UserOrder) TableName() string {
	return "user_orders"
}
