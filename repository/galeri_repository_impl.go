package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type galeriRepositoryImpl struct {
	DB *sql.DB
}

func NewGaleriRepository(db *sql.DB) GaleriRepository {
	return &galeriRepositoryImpl{DB: db}
}

func (r *galeriRepositoryImpl) Create(ctx context.Context, galeri domain.Galeri) (domain.Galeri, error) {
	query := `INSERT INTO galeri (type_galeri_id, title_image, image) VALUES (?, ?, ?)`
	result, err := r.DB.ExecContext(ctx, query, galeri.TypeGaleriID, galeri.TitleImage, galeri.Image)
	if err != nil {
		return galeri, err
	}
	id, _ := result.LastInsertId()
	galeri.ID = int(id)
	return galeri, nil
}

func (r *galeriRepositoryImpl) Update(ctx context.Context, galeri domain.Galeri) (domain.Galeri, error) {
	query := `UPDATE galeri SET type_galeri_id = ?, title_image = ?, image = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, galeri.TypeGaleriID, galeri.TitleImage, galeri.Image, galeri.ID)
	return galeri, err
}

func (r *galeriRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM galeri WHERE id = ?", id)
	return err
}

func (r *galeriRepositoryImpl) FindByID(ctx context.Context, id int) (domain.Galeri, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT id, type_galeri_id, title_image, image FROM galeri WHERE id = ?", id)

	var g domain.Galeri
	err := row.Scan(&g.ID, &g.TypeGaleriID, &g.TitleImage, &g.Image)
	return g, err
}

func (r *galeriRepositoryImpl) FindAll(ctx context.Context) ([]domain.Galeri, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, type_galeri_id, title_image, image FROM galeri")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Galeri
	for rows.Next() {
		var g domain.Galeri
		if err := rows.Scan(&g.ID, &g.TypeGaleriID, &g.TitleImage, &g.Image); err != nil {
			return nil, err
		}
		list = append(list, g)
	}
	return list, nil
}

func (r *galeriRepositoryImpl) FindByUser(ctx context.Context, userID int) ([]domain.Galeri, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, type_galeri_id, title_image, image FROM galeri")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Galeri
	for rows.Next() {
		var g domain.Galeri
		if err := rows.Scan(&g.ID, &g.TypeGaleriID, &g.TitleImage, &g.Image); err != nil {
			return nil, err
		}
		list = append(list, g)
	}
	return list, nil
}
