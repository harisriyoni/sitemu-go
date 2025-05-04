package repository

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	FindByID(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id int) error
}
