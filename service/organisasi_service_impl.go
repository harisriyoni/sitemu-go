package service

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/model/domain"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/repository"
)

type organisasiServiceImpl struct {
	Repo     repository.OrganisasiRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewOrganisasiService(repo repository.OrganisasiRepository, db *sql.DB, validate *validator.Validate) OrganisasiService {
	return &organisasiServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate,
	}
}

func (s *organisasiServiceImpl) CreateOrganisasi(ctx context.Context, userID int, req web.OrganisasiCreateRequest, image multipart.File, imageHeader *multipart.FileHeader) (web.OrganisasiResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	filename, err := helper.SaveUploadedFile(image, imageHeader, "organisasi")
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	org := domain.Organisasi{
		UserID:  userID,
		Jabatan: req.Jabatan,
		Nama:    req.Nama,
		Image:   filename,
	}

	saved, err := s.Repo.Create(ctx, org)
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	return web.OrganisasiResponse{
		ID:      saved.ID,
		UserID:  saved.UserID,
		Jabatan: saved.Jabatan,
		Nama:    saved.Nama,
		Image:   helper.PublicImageURL("organisasi", saved.Image),
	}, nil
}

func (s *organisasiServiceImpl) GetOrganisasiByUserID(ctx context.Context, userID int) ([]web.OrganisasiResponse, error) {
	items, err := s.Repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var results []web.OrganisasiResponse
	for _, item := range items {
		results = append(results, web.OrganisasiResponse{
			ID:      item.ID,
			UserID:  item.UserID,
			Jabatan: item.Jabatan,
			Nama:    item.Nama,
			Image:   helper.PublicImageURL("organisasi", item.Image),
		})
	}
	return results, nil
}

func (s *organisasiServiceImpl) UpdateOrganisasi(
	ctx context.Context,
	id int,
	userID int,
	req web.OrganisasiUpdateRequest,
	image multipart.File,
	imageHeader *multipart.FileHeader,
) (web.OrganisasiResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	org, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	if org.UserID != userID {
		return web.OrganisasiResponse{}, fmt.Errorf("unauthorized access")
	}

	if image != nil {
		// Pakai helper baru: ReplaceUploadedFile
		newFilename, err := helper.ReplaceUploadedFile(org.Image, image, imageHeader, "organisasi")
		if err != nil {
			return web.OrganisasiResponse{}, err
		}
		org.Image = newFilename
	}

	org.Nama = req.Nama
	org.Jabatan = req.Jabatan

	updated, err := s.Repo.Update(ctx, org)
	if err != nil {
		return web.OrganisasiResponse{}, err
	}

	return web.OrganisasiResponse{
		ID:      updated.ID,
		UserID:  updated.UserID,
		Jabatan: updated.Jabatan,
		Nama:    updated.Nama,
		Image:   helper.PublicImageURL("organisasi", updated.Image),
	}, nil
}

func (s *organisasiServiceImpl) DeleteOrganisasi(ctx context.Context, id int, userID int) error {
	org, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if org.UserID != userID { 
		return fmt.Errorf("unauthorized")
	}

	// Hapus file gambar
	err = helper.DeleteFile("organisasi", org.Image)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete image file: %v", err)
	}

	// Hapus dari database
	return s.Repo.Delete(ctx, id)
}

func (s *organisasiServiceImpl) GetAllOrganisasi(ctx context.Context) ([]web.OrganisasiResponse, error) {
	list, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []web.OrganisasiResponse
	for _, item := range list {
		result = append(result, web.OrganisasiResponse{
			ID:      item.ID,
			UserID:  item.UserID,
			Jabatan: item.Jabatan,
			Nama:    item.Nama,
			Image:   helper.PublicImageURL("organisasi", item.Image),
		})
	}
	return result, nil
}
