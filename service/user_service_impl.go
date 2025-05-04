package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/harisriyoni/sitemu-go/helper"
	"github.com/harisriyoni/sitemu-go/model/domain"
	"github.com/harisriyoni/sitemu-go/model/web"
	"github.com/harisriyoni/sitemu-go/repository"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	Repo     repository.UserRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserService(repo repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &userServiceImpl{
		Repo:     repo,
		DB:       db,
		Validate: validate,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req web.UserRegisterRequest) (web.UserResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return web.UserResponse{}, err
	}

	// Cek username apakah sudah digunakan
	_, err = s.Repo.FindByUsername(ctx, req.Username)
	if err == nil {
		return web.UserResponse{}, fmt.Errorf("username '%s' already exists", req.Username)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: string(hashed),
	}

	savedUser, err := s.Repo.Save(ctx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		ID:       savedUser.ID,
		Name:     savedUser.Name,
		Username: savedUser.Username,
	}, nil
}

func (s *userServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) (string, float64, error) {
	user, err := s.Repo.FindByUsername(ctx, req.Username)
	if err != nil {
		return "", 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", 0, err
	}

	token, duration, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return "", 0, err
	}

	return token, duration.Hours(), nil
}

func (s *userServiceImpl) GetProfile(ctx context.Context, id int) (web.UserResponse, error) {
	user, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}, nil
}

func (s *userServiceImpl) UpdateProfile(ctx context.Context, id int, req web.UserUpdateRequest) (web.UserResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		ID:       id,
		Name:     req.Name,
		Username: req.Username,
	}

	updated, err := s.Repo.Update(ctx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		ID:       updated.ID,
		Name:     updated.Name,
		Username: updated.Username,
	}, nil
}

func (s *userServiceImpl) DeleteAccount(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
