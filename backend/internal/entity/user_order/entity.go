package userorder

type (
	UserOrder struct {
		ID          uint
		UserID      uint
		Address     string
		Comment     string
		PhoneNumber string
		Email       string
		Status      OrderStatus
		CreatedAt   int64

		Items []UserOrderItem `gorm:"-"`
	}

	UserOrderItem struct {
		ID          uint
		UserOrderID uint
		ItemID      uint
		Amount      uint // Amount in pieces, e.g. 12 (packets)
	}
)

func (UserOrder) TableName() string {
	return "user_orders"
}

func (UserOrderItem) TableName() string {
	return "user_order_items"
}
