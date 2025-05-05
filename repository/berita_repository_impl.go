package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type beritaRepositoryImpl struct {
	DB *sql.DB
}

func NewBeritaRepository(db *sql.DB) BeritaRepository {
	return &beritaRepositoryImpl{DB: db}
}

func (r *beritaRepositoryImpl) Create(ctx context.Context, berita domain.Berita) (domain.Berita, error) {
	query := `INSERT INTO berita (user_id, title_berita, tanggal, image, deskripsi) VALUES (?, ?, ?, ?, ?)`
	result, err := r.DB.ExecContext(ctx, query, berita.UserID, berita.TitleBerita, berita.Tanggal, berita.Image, berita.Deskripsi)
	if err != nil {
		return berita, err
	}
	id, _ := result.LastInsertId()
	berita.ID = int(id)
	return berita, nil
}

func (r *beritaRepositoryImpl) Update(ctx context.Context, berita domain.Berita) (domain.Berita, error) {
	query := `UPDATE berita SET title_berita = ?, tanggal = ?, image = ?, deskripsi = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, berita.TitleBerita, berita.Tanggal, berita.Image, berita.Deskripsi, berita.ID)
	return berita, err
}

func (r *beritaRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM berita WHERE id = ?", id)
	return err
}

func (r *beritaRepositoryImpl) FindByUserID(ctx context.Context, userID int) ([]domain.Berita, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, title_berita, tanggal, image, deskripsi FROM berita WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Berita
	for rows.Next() {
		var berita domain.Berita
		err := rows.Scan(&berita.ID, &berita.UserID, &berita.TitleBerita, &berita.Tanggal, &berita.Image, &berita.Deskripsi)
		if err != nil {
			return nil, err
		}
		list = append(list, berita)
	}
	return list, nil
}

func (r *beritaRepositoryImpl) FindAll(ctx context.Context) ([]domain.Berita, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, title_berita, tanggal, image, deskripsi FROM berita")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Berita
	for rows.Next() {
		var berita domain.Berita
		err := rows.Scan(&berita.ID, &berita.UserID, &berita.TitleBerita, &berita.Tanggal, &berita.Image, &berita.Deskripsi)
		if err != nil {
			return nil, err
		}
		list = append(list, berita)
	}
	return list, nil
}

func (r *beritaRepositoryImpl) FindByID(ctx context.Context, id int) (domain.Berita, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, user_id, title_berita, tanggal, image, deskripsi FROM berita WHERE id = ?", id)

	var berita domain.Berita
	err := row.Scan(&berita.ID, &berita.UserID, &berita.TitleBerita, &berita.Tanggal, &berita.Image, &berita.Deskripsi)
	return berita, err
}
func (r *beritaRepositoryImpl) FindByUser(ctx context.Context, userID int) ([]domain.Berita, error) {
	var beritas []domain.Berita
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, title_berita, tanggal, image, deskripsi FROM berita WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var berita domain.Berita
		if err := rows.Scan(&berita.ID, &berita.UserID, &berita.TitleBerita, &berita.Tanggal, &berita.Image, &berita.Deskripsi); err != nil {
			return nil, err
		}
		beritas = append(beritas, berita)
	}
	return beritas, err
}
