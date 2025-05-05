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

type galeriServiceImpl struct {
	Repo repository.GaleriRepository
}

func NewGaleriService(repo repository.GaleriRepository) GaleriService {
	return &galeriServiceImpl{Repo: repo}
}

func toImageURL(path string) string {
	cleanPath := strings.TrimPrefix(path, "/")
	cleanPath = strings.ReplaceAll(cleanPath, "\\", "/")
	return fmt.Sprintf("http://localhost:8080/%s", cleanPath)
}

// ✅ CREATE
func (s *galeriServiceImpl) Create(ctx context.Context, request web.GaleriCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error) {
	imagePath := ""
	if imageFile != nil && imageHeader != nil {
		ext := filepath.Ext(imageHeader.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		imagePath = filepath.Join("public", "galeri", filename)

		out, err := os.Create(imagePath)
		if err != nil {
			return web.GaleriResponse{}, err
		}
		defer out.Close()

		_, err = helper.CopyFile(out, imageFile)
		if err != nil {
			return web.GaleriResponse{}, err
		}
	}

	entity := domain.Galeri{
		TitleImage:   request.TitleImage,
		Image:        "/" + imagePath,
		TypeGaleriID: request.TypeGaleriID,
	}

	saved, err := s.Repo.Create(ctx, entity)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	return web.GaleriResponse{
		ID:           saved.ID,
		TitleImage:   saved.TitleImage,
		Image:        toImageURL(saved.Image),
		TypeGaleriID: saved.TypeGaleriID,
	}, nil
}

// ✅ UPDATE
func (s *galeriServiceImpl) Update(ctx context.Context, id int, request web.GaleriUpdateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error) {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	if imageFile != nil && imageHeader != nil {
		if existing.Image != "" {
			_ = os.Remove("." + existing.Image)
		}

		ext := filepath.Ext(imageHeader.Filename)
		filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		newImagePath := filepath.Join("public", "galeri", filename)

		out, err := os.Create(newImagePath)
		if err != nil {
			return web.GaleriResponse{}, err
		}
		defer out.Close()

		_, err = helper.CopyFile(out, imageFile)
		if err != nil {
			return web.GaleriResponse{}, err
		}

		existing.Image = "/" + newImagePath
	}

	existing.TitleImage = request.TitleImage
	existing.TypeGaleriID = request.TypeGaleriID

	updated, err := s.Repo.Update(ctx, existing)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	return web.GaleriResponse{
		ID:           updated.ID,
		TitleImage:   updated.TitleImage,
		Image:        toImageURL(updated.Image),
		TypeGaleriID: updated.TypeGaleriID,
	}, nil
}

// ✅ DELETE
func (s *galeriServiceImpl) Delete(ctx context.Context, id int) error {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existing.Image != "" {
		_ = os.Remove("." + existing.Image)
	}

	return s.Repo.Delete(ctx, id)
}

// ✅ GET ALL
func (s *galeriServiceImpl) GetAll(ctx context.Context) ([]web.GaleriResponse, error) {
	list, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.GaleriResponse
	for _, g := range list {
		responses = append(responses, web.GaleriResponse{
			ID:           g.ID,
			TitleImage:   g.TitleImage,
			Image:        toImageURL(g.Image),
			TypeGaleriID: g.TypeGaleriID,
		})
	}
	return responses, nil
}

// ✅ GET BY ID
func (s *galeriServiceImpl) GetByID(ctx context.Context, id int) (web.GaleriResponse, error) {
	g, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	return web.GaleriResponse{
		ID:           g.ID,
		TitleImage:   g.TitleImage,
		Image:        toImageURL(g.Image),
		TypeGaleriID: g.TypeGaleriID,
	}, nil
}
