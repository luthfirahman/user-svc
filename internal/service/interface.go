package service

import (
	"context"

	"github.com/luthfirahman/user-svc/internal/dto"
	"github.com/luthfirahman/user-svc/internal/repository"
)

type IUserService interface {
	Register(ctx context.Context, r *dto.RegisterRequest) (*repository.User, error)
	Login(ctx context.Context, r dto.LoginRequest) (dto.LoginResponse, error)
	GetMyProfile(ctx context.Context, userID int) (*repository.User, error)
	UpdateMyProfile(ctx context.Context, r *dto.ProfileUpdateRequest, userID int) (*repository.User, error)
}
