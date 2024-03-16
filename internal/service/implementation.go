package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/internal/dto"
	"github.com/luthfirahman/user-svc/internal/repository"
	"github.com/luthfirahman/user-svc/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(ctx context.Context, r *dto.RegisterRequest) (*repository.User, error) {

	user, isExist, _ := s.userRepository.FindByPhoneNumber(ctx, r.PhoneNumber)
	if isExist {
		return user, errors.New("phone number already exists")
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	r.Password = string(hashPassword)

	assignedUser := &repository.User{
		Name:        r.Name,
		PhoneNumber: r.PhoneNumber,
		Password:    r.Password,
	}

	registeredUser, err := s.userRepository.Save(ctx, assignedUser)
	if err != nil {
		return registeredUser, errors.New("failed to create user")
	}

	return registeredUser, nil
}

func (s *userService) Login(ctx context.Context, r dto.LoginRequest) (dto.LoginResponse, error) {

	user, isExist, _ := s.userRepository.FindByPhoneNumber(ctx, r.PhoneNumber)
	if !isExist {
		return dto.LoginResponse{}, errors.New("invalid phone number or password")
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(r.Password))
	if err != nil || !checkPassword {
		return dto.LoginResponse{}, errors.New("invalid phone number or password")
	}

	token := config.NewJWTService().GenerateToken(strconv.Itoa(user.Id), user.PhoneNumber)

	return dto.LoginResponse{
		ID:    user.Id,
		Token: token,
	}, nil

}

func (s userService) GetMyProfile(ctx context.Context, userID int) (*repository.User, error) {
	user, isExist, _ := s.userRepository.FindByID(ctx, userID)
	if !isExist {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) UpdateMyProfile(ctx context.Context, r *dto.ProfileUpdateRequest, userID int) (*repository.User, error) {
	myUser, isExist, _ := s.userRepository.FindByID(ctx, userID)
	if !isExist {
		return myUser, errors.New("user not found")
	}

	existingUser, isExist, _ := s.userRepository.FindByPhoneNumber(ctx, r.PhoneNumber)
	if isExist && existingUser.Id != myUser.Id {
		return &repository.User{}, errors.New("phone number not available")
	}

	assignedUser := &repository.User{
		Id:          myUser.Id,
		Name:        r.Name,
		PhoneNumber: r.PhoneNumber,
	}

	updatedUser, err := s.userRepository.Update(ctx, assignedUser)
	if err != nil {
		return updatedUser, errors.New("failed to update user")
	}

	return updatedUser, nil
}
