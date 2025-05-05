package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/model/domain"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/repository"
)

type beritaServiceImpl struct {
	Repo repository.BeritaRepository
}

func NewBeritaService(repo repository.BeritaRepository) BeritaService {
	return &beritaServiceImpl{Repo: repo}
}

const beritaFolderID = "YOUR_BERITA_FOLDER_ID"

func (s *beritaServiceImpl) Create(ctx context.Context, userID int, request web.BeritaCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.BeritaResponse, error) {
	imageID := ""
	if imageFile != nil && imageHeader != nil {
		driveID, _, err := helper.UploadToDrive(imageFile, imageHeader, beritaFolderID)
		if err != nil {
			return web.BeritaResponse{}, err
		}
		imageID = driveID
	}

	berita := domain.Berita{
		TitleBerita: request.TitleBerita,
		Tanggal:     request.Tanggal,
		Deskripsi:   request.Deskripsi,
		Image:       imageID,
		UserID:      userID,
	}

	saved, err := s.Repo.Create(ctx, berita)
	if err != nil {
		return web.BeritaResponse{}, err
	}

	return web.BeritaResponse{
		ID:          saved.ID,
		TitleBerita: saved.TitleBerita,
		Tanggal:     saved.Tanggal,
		Deskripsi:   saved.Deskripsi,
		Image:       helper.PublicImageURLDrive(saved.Image),
		UserID:      saved.UserID,
	}, nil
}

func (s *beritaServiceImpl) Update(ctx context.Context, id int, userID int, request web.BeritaUpdateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.BeritaResponse, error) {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.BeritaResponse{}, err
	}
	if existing.UserID != userID {
		return web.BeritaResponse{}, fmt.Errorf("unauthorized update")
	}

	if imageFile != nil && imageHeader != nil {
		if existing.Image != "" {
			_ = helper.DeleteFromDrive(existing.Image)
		}
		driveID, _, err := helper.UploadToDrive(imageFile, imageHeader, beritaFolderID)
		if err != nil {
			return web.BeritaResponse{}, err
		}
		existing.Image = driveID
	}

	existing.TitleBerita = request.TitleBerita
	existing.Tanggal = request.Tanggal
	existing.Deskripsi = request.Deskripsi

	updated, err := s.Repo.Update(ctx, existing)
	if err != nil {
		return web.BeritaResponse{}, err
	}

	return web.BeritaResponse{
		ID:          updated.ID,
		TitleBerita: updated.TitleBerita,
		Tanggal:     updated.Tanggal,
		Deskripsi:   updated.Deskripsi,
		Image:       helper.PublicImageURLDrive(updated.Image),
		UserID:      updated.UserID,
	}, nil
}

func (s *beritaServiceImpl) Delete(ctx context.Context, id int, userID int) error {
	berita, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if berita.UserID != userID {
		return fmt.Errorf("unauthorized delete")
	}

	if berita.Image != "" {
		_ = helper.DeleteFromDrive(berita.Image)
	}

	return s.Repo.Delete(ctx, id)
}

func (s *beritaServiceImpl) GetAll(ctx context.Context) ([]web.BeritaResponse, error) {
	beritas, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.BeritaResponse
	for _, b := range beritas {
		responses = append(responses, web.BeritaResponse{
			ID:          b.ID,
			TitleBerita: b.TitleBerita,
			Tanggal:     b.Tanggal,
			Image:       helper.PublicImageURLDrive(b.Image),
			Deskripsi:   b.Deskripsi,
			UserID:      b.UserID,
		})
	}

	return responses, nil
}

func (s *beritaServiceImpl) GetByID(ctx context.Context, id int) (web.BeritaResponse, error) {
	b, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.BeritaResponse{}, err
	}

	return web.BeritaResponse{
		ID:          b.ID,
		TitleBerita: b.TitleBerita,
		Tanggal:     b.Tanggal,
		Image:       helper.PublicImageURLDrive(b.Image),
		Deskripsi:   b.Deskripsi,
		UserID:      b.UserID,
	}, nil
}

func (s *beritaServiceImpl) GetByUser(ctx context.Context, userID int) ([]web.BeritaResponse, error) {
	beritas, err := s.Repo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []web.BeritaResponse
	for _, b := range beritas {
		responses = append(responses, web.BeritaResponse{
			ID:          b.ID,
			TitleBerita: b.TitleBerita,
			Tanggal:     b.Tanggal,
			Image:       helper.PublicImageURLDrive(b.Image),
			Deskripsi:   b.Deskripsi,
			UserID:      b.UserID,
		})
	}
	return responses, nil
}
