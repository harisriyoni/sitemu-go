package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type prestasiRepositoryImpl struct {
	DB *sql.DB
}

func NewPrestasiRepository(db *sql.DB) PrestasiRepository {
	return &prestasiRepositoryImpl{DB: db}
}

func (r *prestasiRepositoryImpl) Create(ctx context.Context, p domain.Prestasi) (domain.Prestasi, error) {
	query := `INSERT INTO prestasi (user_id, title, tahun, prestasi, deskripsi) VALUES (?, ?, ?, ?, ?)`
	result, err := r.DB.ExecContext(ctx, query, p.UserID, p.Title, p.Tahun, p.Prestasi, p.Deskripsi)
	if err != nil {
		return p, err
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return p, nil
}

func (r *prestasiRepositoryImpl) Update(ctx context.Context, p domain.Prestasi) (domain.Prestasi, error) {
	query := `UPDATE prestasi SET title = ?, tahun = ?, prestasi = ?, deskripsi = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, p.Title, p.Tahun, p.Prestasi, p.Deskripsi, p.ID)
	return p, err
}

func (r *prestasiRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM prestasi WHERE id = ?", id)
	return err
}

func (r *prestasiRepositoryImpl) FindByID(ctx context.Context, id int) (domain.Prestasi, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, user_id, title, tahun, prestasi, deskripsi FROM prestasi WHERE id = ?", id)

	var p domain.Prestasi
	err := row.Scan(&p.ID, &p.UserID, &p.Title, &p.Tahun, &p.Prestasi, &p.Deskripsi)
	return p, err
}

func (r *prestasiRepositoryImpl) FindByUser(ctx context.Context, userID int) ([]domain.Prestasi, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, title, tahun, prestasi, deskripsi FROM prestasi WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Prestasi
	for rows.Next() {
		var p domain.Prestasi
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Tahun, &p.Prestasi, &p.Deskripsi)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func (r *prestasiRepositoryImpl) FindAll(ctx context.Context) ([]domain.Prestasi, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, title, tahun, prestasi, deskripsi FROM prestasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Prestasi
	for rows.Next() {
		var p domain.Prestasi
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Tahun, &p.Prestasi, &p.Deskripsi)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}
