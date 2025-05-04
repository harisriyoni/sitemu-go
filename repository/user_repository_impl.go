package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Save(ctx context.Context, user domain.User) (domain.User, error) {
	stmt := `INSERT INTO users (name, username, password, created_at, updated_at)
	         VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	res, err := r.DB.ExecContext(ctx, stmt, user.Name, user.Username, user.Password)
	if err != nil {
		return user, err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `SELECT id, name, username, password, created_at, updated_at FROM users WHERE username = ?`
	row := r.DB.QueryRowContext(ctx, query, username)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int) (domain.User, error) {
	query := `SELECT id, name, username, password, created_at, updated_at FROM users WHERE id = ?`
	row := r.DB.QueryRowContext(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *userRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	stmt := `UPDATE users SET name=?, username=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := r.DB.ExecContext(ctx, stmt, user.Name, user.Username, user.ID)
	return user, err
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)
	return err
}
