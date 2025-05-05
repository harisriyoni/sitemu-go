package service

import (
	"context"
	"mime/multipart"

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

const galeriFolderID = "YOUR_GALERI_FOLDER_ID"

func (s *galeriServiceImpl) Create(ctx context.Context, request web.GaleriCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error) {
	imageID := ""
	if imageFile != nil && imageHeader != nil {
		driveID, _, err := helper.UploadToDrive(imageFile, imageHeader, galeriFolderID)
		if err != nil {
			return web.GaleriResponse{}, err
		}
		imageID = driveID
	}

	entity := domain.Galeri{
		TitleImage:   request.TitleImage,
		Image:        imageID,
		TypeGaleriID: request.TypeGaleriID,
	}

	saved, err := s.Repo.Create(ctx, entity)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	return web.GaleriResponse{
		ID:           saved.ID,
		TitleImage:   saved.TitleImage,
		Image:        helper.PublicImageURLDrive(saved.Image),
		TypeGaleriID: saved.TypeGaleriID,
	}, nil
}

func (s *galeriServiceImpl) Update(ctx context.Context, id int, request web.GaleriUpdateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error) {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	if imageFile != nil && imageHeader != nil {
		if existing.Image != "" {
			_ = helper.DeleteFromDrive(existing.Image)
		}
		driveID, _, err := helper.UploadToDrive(imageFile, imageHeader, galeriFolderID)
		if err != nil {
			return web.GaleriResponse{}, err
		}
		existing.Image = driveID
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
		Image:        helper.PublicImageURLDrive(updated.Image),
		TypeGaleriID: updated.TypeGaleriID,
	}, nil
}

func (s *galeriServiceImpl) Delete(ctx context.Context, id int) error {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existing.Image != "" {
		_ = helper.DeleteFromDrive(existing.Image)
	}

	return s.Repo.Delete(ctx, id)
}

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
			Image:        helper.PublicImageURLDrive(g.Image),
			TypeGaleriID: g.TypeGaleriID,
		})
	}
	return responses, nil
}

func (s *galeriServiceImpl) GetByID(ctx context.Context, id int) (web.GaleriResponse, error) {
	g, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.GaleriResponse{}, err
	}

	return web.GaleriResponse{
		ID:           g.ID,
		TitleImage:   g.TitleImage,
		Image:        helper.PublicImageURLDrive(g.Image),
		TypeGaleriID: g.TypeGaleriID,
	}, nil
}
