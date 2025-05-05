package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type PrestasiRepository interface {
	Create(ctx context.Context, prestasi domain.Prestasi) (domain.Prestasi, error)
	Update(ctx context.Context, prestasi domain.Prestasi) (domain.Prestasi, error)
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (domain.Prestasi, error)
	FindByUser(ctx context.Context, userID int) ([]domain.Prestasi, error)
	FindAll(ctx context.Context) ([]domain.Prestasi, error)
}
