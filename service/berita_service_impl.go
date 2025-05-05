package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
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

func toFullImageURL(path string) string {
	cleanPath := strings.TrimPrefix(path, "/")
	cleanPath = strings.ReplaceAll(cleanPath, "\\", "/")
	return fmt.Sprintf("http://localhost:8080/%s", cleanPath)
}

func (s *beritaServiceImpl) Create(ctx context.Context, userID int, request web.BeritaCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.BeritaResponse, error) {
	imagePath := ""
	if imageFile != nil && imageHeader != nil {
		ext := filepath.Ext(imageHeader.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		imagePath = filepath.Join("public", "berita", filename)

		out, err := os.Create(imagePath)
		if err != nil {
			return web.BeritaResponse{}, err
		}
		defer out.Close()

		_, err = helper.CopyFile(out, imageFile)
		if err != nil {
			return web.BeritaResponse{}, err
		}
	}

	berita := domain.Berita{
		TitleBerita: request.TitleBerita,
		Tanggal:     request.Tanggal,
		Deskripsi:   request.Deskripsi,
		Image:       "/" + imagePath,
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
		Image:       toFullImageURL(saved.Image),
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
			os.Remove("." + existing.Image)
		}

		ext := filepath.Ext(imageHeader.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		newImagePath := filepath.Join("public", "berita", filename)

		out, err := os.Create(newImagePath)
		if err != nil {
			return web.BeritaResponse{}, err
		}
		defer out.Close()

		_, err = helper.CopyFile(out, imageFile)
		if err != nil {
			return web.BeritaResponse{}, err
		}

		existing.Image = "/" + newImagePath
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
		Image:       toFullImageURL(updated.Image),
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
		os.Remove("." + berita.Image)
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
			Image:       toFullImageURL(b.Image),
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
		Image:       toFullImageURL(b.Image),
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
			Image:       toFullImageURL(b.Image),
			Deskripsi:   b.Deskripsi,
			UserID:      b.UserID,
		})
	}
	return responses, nil
}
