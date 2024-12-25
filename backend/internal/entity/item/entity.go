package item

type Item struct {
	ID          uint
	Title       string
	Description string
	CreatedAt   int64
}

func (Item) TableName() string {
	return "items"
}
