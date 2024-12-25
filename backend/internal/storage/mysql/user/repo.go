package user

import (
	"backend/internal/entity/user"
	"backend/pkg/db"
	"context"
)

type repo struct {
	db db.Database
}

func NewRepo(db db.Database) repo {
	return repo{db}
}

func (r repo) GetUserByID(ctx context.Context, userID uint) (user user.User, err error) {
	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM users WHERE id = $1`, userID).Scan(&user).Error
	return
}

func (r repo) GetUserByEmail(ctx context.Context, email string) (user user.User, err error) {
	err = r.db.Instance().WithContext(ctx).Raw(`
		SELECT * FROM users WHERE email = $1`, email).Scan(&user).Error
	return
}

func (r repo) CreateUser(ctx context.Context, body user.User) (user.User, error) {
	err := r.db.Instance().WithContext(ctx).Create(&body).Error
	return body, err
}

func (r repo) UpdateUser(ctx context.Context, body user.User) error {
	return r.db.Instance().WithContext(ctx).Where("id", body.ID).Updates(&body).Error
}

func (r repo) DeleteUser(ctx context.Context, userID uint) error {
	return r.db.Instance().WithContext(ctx).Exec(`
		DELETE FROM users WHERE id = $1`, userID).Error
}
