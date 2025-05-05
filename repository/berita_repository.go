package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type BeritaRepository interface {
	Create(ctx context.Context, berita domain.Berita) (domain.Berita, error)
	Update(ctx context.Context, berita domain.Berita) (domain.Berita, error)
	Delete(ctx context.Context, id int) error
	FindByUserID(ctx context.Context, userID int) ([]domain.Berita, error)
	FindAll(ctx context.Context) ([]domain.Berita, error)
	FindByID(ctx context.Context, id int) (domain.Berita, error)
	FindByUser(ctx context.Context, userID int) ([]domain.Berita, error)
}
