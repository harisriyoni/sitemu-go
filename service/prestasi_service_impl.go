package service

import (
	"context"
	"fmt"

	"github.com/harisriyoni/sitemu-go/model/domain"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/repository"
)

type prestasiServiceImpl struct {
	Repo repository.PrestasiRepository
}

func NewPrestasiService(repo repository.PrestasiRepository) PrestasiService {
	return &prestasiServiceImpl{Repo: repo}
}

func (s *prestasiServiceImpl) Create(ctx context.Context, userID int, request web.PrestasiCreateRequest) (web.PrestasiResponse, error) {
	entity := domain.Prestasi{
		UserID:    userID,
		Title:     request.Title,
		Tahun:     request.Tahun,
		Prestasi:  request.Prestasi,
		Deskripsi: request.Deskripsi,
	}

	result, err := s.Repo.Create(ctx, entity)
	if err != nil {
		return web.PrestasiResponse{}, err
	}

	return toPrestasiResponse(result), nil
}

func (s *prestasiServiceImpl) Update(ctx context.Context, id int, userID int, request web.PrestasiUpdateRequest) (web.PrestasiResponse, error) {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.PrestasiResponse{}, err
	}
	if existing.UserID != userID {
		return web.PrestasiResponse{}, fmt.Errorf("unauthorized update")
	}

	existing.Title = request.Title
	existing.Tahun = request.Tahun
	existing.Prestasi = request.Prestasi
	existing.Deskripsi = request.Deskripsi

	updated, err := s.Repo.Update(ctx, existing)
	if err != nil {
		return web.PrestasiResponse{}, err
	}

	return toPrestasiResponse(updated), nil
}

func (s *prestasiServiceImpl) Delete(ctx context.Context, id int, userID int) error {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existing.UserID != userID {
		return fmt.Errorf("unauthorized delete")
	}

	return s.Repo.Delete(ctx, id)
}

func (s *prestasiServiceImpl) GetAll(ctx context.Context) ([]web.PrestasiResponse, error) {
	list, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.PrestasiResponse
	for _, item := range list {
		responses = append(responses, toPrestasiResponse(item))
	}
	return responses, nil
}

func (s *prestasiServiceImpl) GetByID(ctx context.Context, id int) (web.PrestasiResponse, error) {
	data, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.PrestasiResponse{}, err
	}
	return toPrestasiResponse(data), nil
}

func (s *prestasiServiceImpl) GetByUser(ctx context.Context, userID int) ([]web.PrestasiResponse, error) {
	list, err := s.Repo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []web.PrestasiResponse
	for _, item := range list {
		responses = append(responses, toPrestasiResponse(item))
	}
	return responses, nil
}

func toPrestasiResponse(p domain.Prestasi) web.PrestasiResponse {
	return web.PrestasiResponse{
		ID:        p.ID,
		UserID:    p.UserID,
		Title:     p.Title,
		Tahun:     p.Tahun,
		Prestasi:  p.Prestasi,
		Deskripsi: p.Deskripsi,
	}
}
