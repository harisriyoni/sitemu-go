package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type OrganisasiRepository interface {
	Create(ctx context.Context, organisasi domain.Organisasi) (domain.Organisasi, error)
	FindByUserID(ctx context.Context, userID int) ([]domain.Organisasi, error)
	GetAll(ctx context.Context) ([]domain.Organisasi, error)
	FindByID(ctx context.Context, id int) (domain.Organisasi, error)
	Update(ctx context.Context, organisasi domain.Organisasi) (domain.Organisasi, error)
	Delete(ctx context.Context, id int) error
}
