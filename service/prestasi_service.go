package service

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type PrestasiService interface {
	Create(ctx context.Context, userID int, request web.PrestasiCreateRequest) (web.PrestasiResponse, error)
	Update(ctx context.Context, id int, userID int, request web.PrestasiUpdateRequest) (web.PrestasiResponse, error)
	Delete(ctx context.Context, id int, userID int) error

	GetAll(ctx context.Context) ([]web.PrestasiResponse, error)
	GetByID(ctx context.Context, id int) (web.PrestasiResponse, error)
	GetByUser(ctx context.Context, userID int) ([]web.PrestasiResponse, error)
}
