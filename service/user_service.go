package service

import (
	"context"

	"github.com/harisriyoni/sitemu-go/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.UserLoginRequest) (string, float64, error)
	GetProfile(ctx context.Context, id int) (web.UserResponse, error)
	UpdateProfile(ctx context.Context, id int, req web.UserUpdateRequest) (web.UserResponse, error)
	DeleteAccount(ctx context.Context, id int) error
}
