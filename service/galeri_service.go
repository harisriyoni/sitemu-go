package service

import (
	"context"
	"mime/multipart"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type GaleriService interface {
	Create(ctx context.Context, request web.GaleriCreateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error)
	Update(ctx context.Context, id int, request web.GaleriUpdateRequest, imageFile multipart.File, imageHeader *multipart.FileHeader) (web.GaleriResponse, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]web.GaleriResponse, error)
	GetByID(ctx context.Context, id int) (web.GaleriResponse, error)
}
