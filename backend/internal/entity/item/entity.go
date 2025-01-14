package item

type Item struct {
	ID          uint
	Title       string
	Description string
	Price       int64 // Price in cents
	ImageURL    string
	CreatedAt   int64
}

func (Item) TableName() string {
	return "items"
}
