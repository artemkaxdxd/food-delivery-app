package item

import (
	"backend/internal/entity/item"
	"time"
)

type (
	Item struct {
		Title       string `json:"title" valid:"stringlength(3|100)"`
		Description string `json:"description"`
	}
)

func (i Item) ToEntity() item.Item {
	return item.Item{
		Title:       i.Title,
		Description: i.Description,
		CreatedAt:   time.Now().Unix(),
	}
}

func (i Item) ToEntityUpdate(itemID uint) item.Item {
	return item.Item{
		ID:          itemID,
		Title:       i.Title,
		Description: i.Description,
	}
}
