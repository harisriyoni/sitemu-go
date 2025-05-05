package service

import (
	"context"
	"mime/multipart"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type BeritaService interface {
	Create(ctx context.Context, userID int, request web.BeritaCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.BeritaResponse, error)
	Update(ctx context.Context, id int, userID int, request web.BeritaUpdateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.BeritaResponse, error)
	Delete(ctx context.Context, id int, userID int) error
	GetAll(ctx context.Context) ([]web.BeritaResponse, error)
	GetByID(ctx context.Context, id int) (web.BeritaResponse, error)
	GetByUser(ctx context.Context, userID int) ([]web.BeritaResponse, error)
}
