package service

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type TypeGaleriService interface {
	Create(ctx context.Context, request web.TypeGaleriCreateRequest, userID int) (web.TypeGaleriResponse, error)
	Update(ctx context.Context, id int, userID int, request web.TypeGaleriUpdateRequest) (web.TypeGaleriResponse, error)
	Delete(ctx context.Context, id int, userID int) error

	GetAll(ctx context.Context) ([]web.TypeGaleriResponse, error)                // ✅ Public
	GetByID(ctx context.Context, id int) (web.TypeGaleriResponse, error)         // ✅ Public
	GetByUser(ctx context.Context, userID int) ([]web.TypeGaleriResponse, error) // ✅ Protected
}
