package user

import (
	"backend/internal/entity/item"
	"backend/pkg/logger"
	"context"
)

type (
	repo interface {
		GetItems(ctx context.Context, search string, limit, offset uint) ([]item.Item, error)
		CreateItem(ctx context.Context, body item.Item) (item.Item, error)
		UpdateItem(ctx context.Context, body item.Item) error
		DeleteItem(ctx context.Context, itemID uint) error
	}

	service struct {
		repo repo
		l    logger.Logger
	}
)

func NewService(repo repo, l logger.Logger) service {
	return service{repo, l}
}
