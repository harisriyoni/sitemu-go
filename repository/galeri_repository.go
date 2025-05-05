package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type GaleriRepository interface {
	Create(ctx context.Context, galeri domain.Galeri) (domain.Galeri, error)
	Update(ctx context.Context, galeri domain.Galeri) (domain.Galeri, error)
	Delete(ctx context.Context, id int) error

	FindByID(ctx context.Context, id int) (domain.Galeri, error)
	FindAll(ctx context.Context) ([]domain.Galeri, error)
	FindByUser(ctx context.Context, userID int) ([]domain.Galeri, error)
}
