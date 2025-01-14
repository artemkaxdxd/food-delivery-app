package item

import "backend/internal/entity/item"

type Item struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	CreatedAt   int64  `json:"created_at"`
}

func ItemToResponse(item item.Item) Item {
	return Item(item)
}

func ItemsToResponse(items []item.Item) []Item {
	if len(items) == 0 {
		return nil
	}

	resp := make([]Item, len(items))
	for i, v := range items {
		resp[i] = ItemToResponse(v)
	}
	return resp
}
