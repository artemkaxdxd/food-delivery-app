package item

import (
	"backend/internal/entity/item"
	"backend/pkg/db"
	"context"
	"fmt"
)

type repo struct {
	db db.Database
}

func NewRepo(db db.Database) repo {
	return repo{db}
}

func (r repo) GetItems(ctx context.Context, search string, limit, offset uint) (items []item.Item, err error) {
	search = fmt.Sprintf("%%%s%%", search)

	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM items 
		WHERE title LIKE $1
		LIMIT $2 OFFSET $3`, search, limit, offset).Scan(&items).Error
	return
}

func (r repo) CreateItem(ctx context.Context, body item.Item) (item.Item, error) {
	err := r.db.Instance().WithContext(ctx).Create(&body).Error
	return body, err
}

func (r repo) UpdateItem(ctx context.Context, body item.Item) error {
	return r.db.Instance().WithContext(ctx).Where("id", body.ID).Updates(&body).Error
}

func (r repo) DeleteItem(ctx context.Context, itemID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM items WHERE id = $1`, itemID).Error
}
