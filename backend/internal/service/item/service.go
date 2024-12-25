package user

import (
	"backend/config"
	request "backend/internal/controllers/http/request/item"
	response "backend/internal/controllers/http/response/item"
	"context"
)

func (s service) GetItems(ctx context.Context, search string, limit, offset uint) ([]response.Item, config.ServiceCode, error) {
	items, err := s.repo.GetItems(ctx, search, limit, offset)
	return response.ItemsToResponse(items), config.DBErrToServiceCode(err), err
}

func (s service) CreateItem(ctx context.Context, body request.Item) (response.Item, config.ServiceCode, error) {
	item, err := s.repo.CreateItem(ctx, body.ToEntity())
	return response.ItemToResponse(item), config.DBErrToServiceCode(err), err
}

func (s service) UpdateItem(ctx context.Context, itemID uint, body request.Item) (config.ServiceCode, error) {
	err := s.repo.UpdateItem(ctx, body.ToEntityUpdate(itemID))
	return config.DBErrToServiceCode(err), err
}

func (s service) DeleteItem(ctx context.Context, itemID uint) (config.ServiceCode, error) {
	err := s.repo.DeleteItem(ctx, itemID)
	return config.DBErrToServiceCode(err), err
}
