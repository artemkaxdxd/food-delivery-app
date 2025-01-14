package item

import (
	"backend/internal/entity/item"
	"time"
)

type (
	Item struct {
		Title       string `json:"title" valid:"stringlength(3|100)"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
		ImageURL    string `json:"image_url"`
	}
)

func (i Item) ToEntity() item.Item {
	return item.Item{
		Title:       i.Title,
		Description: i.Description,
		Price:       i.Price,
		ImageURL:    i.ImageURL,
		CreatedAt:   time.Now().Unix(),
	}
}

func (i Item) ToEntityUpdate(itemID uint) item.Item {
	return item.Item{
		ID:          itemID,
		Title:       i.Title,
		Description: i.Description,
		Price:       i.Price,
		ImageURL:    i.ImageURL,
	}
}
