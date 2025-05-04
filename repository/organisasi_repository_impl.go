package repository

import (
	"context"
	"database/sql"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type organisasiRepositoryImpl struct {
	DB *sql.DB
}

func NewOrganisasiRepository(db *sql.DB) OrganisasiRepository {
	return &organisasiRepositoryImpl{DB: db}
}

func (r *organisasiRepositoryImpl) Create(ctx context.Context, organisasi domain.Organisasi) (domain.Organisasi, error) {
	stmt := `INSERT INTO organisasi (user_id, jabatan, nama, image) VALUES (?, ?, ?, ?)`
	res, err := r.DB.ExecContext(ctx, stmt, organisasi.UserID, organisasi.Jabatan, organisasi.Nama, organisasi.Image)
	if err != nil {
		return organisasi, err
	}
	id, _ := res.LastInsertId()
	organisasi.ID = int(id)
	return organisasi, nil
}

func (r *organisasiRepositoryImpl) FindByUserID(ctx context.Context, userID int) ([]domain.Organisasi, error) {
	query := `SELECT id, user_id, jabatan, nama, image FROM organisasi WHERE user_id = ?`
	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Organisasi
	for rows.Next() {
		var org domain.Organisasi
		err := rows.Scan(&org.ID, &org.UserID, &org.Jabatan, &org.Nama, &org.Image)
		if err != nil {
			return nil, err
		}
		list = append(list, org)
	}
	return list, nil
}

func (r *organisasiRepositoryImpl) FindByID(ctx context.Context, id int) (domain.Organisasi, error) {
	query := `SELECT id, user_id, jabatan, nama, image FROM organisasi WHERE id = ?`
	row := r.DB.QueryRowContext(ctx, query, id)

	var org domain.Organisasi
	err := row.Scan(&org.ID, &org.UserID, &org.Jabatan, &org.Nama, &org.Image)
	return org, err
}

func (r *organisasiRepositoryImpl) Update(ctx context.Context, organisasi domain.Organisasi) (domain.Organisasi, error) {
	stmt := `UPDATE organisasi SET jabatan = ?, nama = ?, image = ? WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, stmt, organisasi.Jabatan, organisasi.Nama, organisasi.Image, organisasi.ID)
	return organisasi, err
}

func (r *organisasiRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM organisasi WHERE id = ?`, id)
	return err
}
func (r *organisasiRepositoryImpl) GetAll(ctx context.Context) ([]domain.Organisasi, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, user_id, jabatan, nama, image FROM organisasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Organisasi
	for rows.Next() {
		var org domain.Organisasi
		err := rows.Scan(&org.ID, &org.UserID, &org.Jabatan, &org.Nama, &org.Image)
		if err != nil {
			return nil, err
		}
		list = append(list, org)
	}

	return list, nil
}
