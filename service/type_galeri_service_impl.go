package service

import (
	"context"
	"fmt"

	"github.com/harisriyoni/sitemu-go/model/domain"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/repository"
)

type typeGaleriServiceImpl struct {
	Repo repository.TypeGaleriRepository
}

func NewTypeGaleriService(repo repository.TypeGaleriRepository) TypeGaleriService {
	return &typeGaleriServiceImpl{Repo: repo}
}

// Create TypeGaleri
func (s *typeGaleriServiceImpl) Create(ctx context.Context, request web.TypeGaleriCreateRequest, userID int) (web.TypeGaleriResponse, error) {
	data := domain.TypeGaleri{
		Type:   request.Type,
		UserID: userID,
	}
	saved, err := s.Repo.Create(ctx, data)
	if err != nil {
		return web.TypeGaleriResponse{}, err
	}
	return web.TypeGaleriResponse{
		ID:     saved.ID,
		Type:   saved.Type,
		UserID: saved.UserID,
	}, nil
}

// Update TypeGaleri
func (s *typeGaleriServiceImpl) Update(ctx context.Context, id int, userID int, request web.TypeGaleriUpdateRequest) (web.TypeGaleriResponse, error) {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.TypeGaleriResponse{}, err
	}
	if existing.UserID != userID {
		return web.TypeGaleriResponse{}, fmt.Errorf("unauthorized update")
	}

	existing.Type = request.Type

	updated, err := s.Repo.Update(ctx, existing)
	if err != nil {
		return web.TypeGaleriResponse{}, err
	}

	return web.TypeGaleriResponse{
		ID:     updated.ID,
		Type:   updated.Type,
		UserID: updated.UserID,
	}, nil
}

// Delete TypeGaleri
func (s *typeGaleriServiceImpl) Delete(ctx context.Context, id int, userID int) error {
	existing, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existing.UserID != userID {
		return fmt.Errorf("unauthorized delete")
	}

	return s.Repo.Delete(ctx, id)
}

// GetAll - Public
func (s *typeGaleriServiceImpl) GetAll(ctx context.Context) ([]web.TypeGaleriResponse, error) {
	list, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []web.TypeGaleriResponse
	for _, t := range list {
		result = append(result, web.TypeGaleriResponse{
			ID:     t.ID,
			Type:   t.Type,
			UserID: t.UserID,
		})
	}
	return result, nil
}

// GetByID - Public
func (s *typeGaleriServiceImpl) GetByID(ctx context.Context, id int) (web.TypeGaleriResponse, error) {
	t, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.TypeGaleriResponse{}, err
	}
	return web.TypeGaleriResponse{
		ID:     t.ID,
		Type:   t.Type,
		UserID: t.UserID,
	}, nil
}

// GetByUser - Protected
func (s *typeGaleriServiceImpl) GetByUser(ctx context.Context, userID int) ([]web.TypeGaleriResponse, error) {
	list, err := s.Repo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []web.TypeGaleriResponse
	for _, t := range list {
		result = append(result, web.TypeGaleriResponse{
			ID:     t.ID,
			Type:   t.Type,
			UserID: t.UserID,
		})
	}
	return result, nil
}
