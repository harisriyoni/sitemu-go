package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type typeGaleriRepositoryImpl struct {
	DB *sql.DB
}

func NewTypeGaleriRepository(db *sql.DB) TypeGaleriRepository {
	return &typeGaleriRepositoryImpl{DB: db}
}

func (r *typeGaleriRepositoryImpl) Create(ctx context.Context, galeri domain.TypeGaleri) (domain.TypeGaleri, error) {
	query := `INSERT INTO type_galeri (type, user_id) VALUES (?, ?)`
	result, err := r.DB.ExecContext(ctx, query, galeri.Type, galeri.UserID)
	if err != nil {
		return galeri, err
	}
	id, _ := result.LastInsertId()
	galeri.ID = int(id)
	return galeri, nil
}

func (r *typeGaleriRepositoryImpl) Update(ctx context.Context, galeri domain.TypeGaleri) (domain.TypeGaleri, error) {
	query := `UPDATE type_galeri SET type = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, galeri.Type, galeri.ID)
	return galeri, err
}

func (r *typeGaleriRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM type_galeri WHERE id = ?", id)
	return err
}

func (r *typeGaleriRepositoryImpl) FindByUser(ctx context.Context, userID int) ([]domain.TypeGaleri, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, type, user_id FROM type_galeri WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.TypeGaleri
	for rows.Next() {
		var item domain.TypeGaleri
		err := rows.Scan(&item.ID, &item.Type, &item.UserID)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *typeGaleriRepositoryImpl) FindAll(ctx context.Context) ([]domain.TypeGaleri, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, type, user_id FROM type_galeri")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.TypeGaleri
	for rows.Next() {
		var item domain.TypeGaleri
		err := rows.Scan(&item.ID, &item.Type, &item.UserID)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *typeGaleriRepositoryImpl) FindByID(ctx context.Context, id int) (domain.TypeGaleri, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, type, user_id FROM type_galeri WHERE id = ?", id)

	var item domain.TypeGaleri
	err := row.Scan(&item.ID, &item.Type, &item.UserID)
	return item, err
}
