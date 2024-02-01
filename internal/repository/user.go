package repository

import (
	"context"
	"database/sql"
	"test/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create は新しいユーザーをデータベースに追加します。
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", u.Name, u.Email)
	return err
}

// Read はユーザー情報を取得します。
func (r *UserRepository) Read(ctx context.Context, id int) (*model.User, error) {
	u := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Update はユーザー情報を更新します。
func (r *UserRepository) Update(ctx context.Context, u *model.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
	return err
}

// Delete はユーザーを削除します。
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
