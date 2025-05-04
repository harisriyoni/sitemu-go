package service

import (
	"context"
	"mime/multipart"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type OrganisasiService interface {
	CreateOrganisasi(ctx context.Context, userID int, req web.OrganisasiCreateRequest, image multipart.File, imageHeader *multipart.FileHeader) (web.OrganisasiResponse, error)
	GetOrganisasiByUserID(ctx context.Context, userID int) ([]web.OrganisasiResponse, error)
	UpdateOrganisasi(ctx context.Context, id int, userID int, req web.OrganisasiUpdateRequest, image multipart.File, imageHeader *multipart.FileHeader) (web.OrganisasiResponse, error)
	DeleteOrganisasi(ctx context.Context, id int, userID int) error
	GetAllOrganisasi(ctx context.Context) ([]web.OrganisasiResponse, error)
}
