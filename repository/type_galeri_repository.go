package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type TypeGaleriRepository interface {
	Create(ctx context.Context, galeri domain.TypeGaleri) (domain.TypeGaleri, error)
	Update(ctx context.Context, galeri domain.TypeGaleri) (domain.TypeGaleri, error)
	Delete(ctx context.Context, id int) error
	FindByUser(ctx context.Context, userID int) ([]domain.TypeGaleri, error)
	FindAll(ctx context.Context) ([]domain.TypeGaleri, error)
	FindByID(ctx context.Context, id int) (domain.TypeGaleri, error)
}
