package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/luthfirahman/user-svc/internal/dto"
	"github.com/luthfirahman/user-svc/internal/repository"
	"github.com/luthfirahman/user-svc/internal/repository/mocks"
	"github.com/luthfirahman/user-svc/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetMyProfile(t *testing.T) {
	ctx := context.Background()
	mockUserRepo := &mocks.IUserRepository{}

	t.Run("Success to get profile", func(t *testing.T) {
		expectedUser1 := &repository.User{Id: 1, Name: "John Doe", PhoneNumber: "+6281234567890"}
		mockUserRepo.On("FindByID", ctx, 1).Return(expectedUser1, true, nil)

		svc := service.NewUserService(mockUserRepo)
		user, err := svc.GetMyProfile(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser1, user)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Failed to get profile", func(t *testing.T) {
		expectedUser2 := &repository.User{}
		expectedErr := errors.New("repository error")
		mockUserRepo.On("FindByID", ctx, 2).Return(expectedUser2, false, expectedErr)

		svc := service.NewUserService(mockUserRepo)
		user, err := svc.GetMyProfile(ctx, 2)

		assert.Error(t, err)
		assert.Equal(t, expectedUser2, user)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdateMyProfile(t *testing.T) {
	ctx := context.Background()
	mockUserRepo := &mocks.IUserRepository{}

	t.Run("Success to update profile", func(t *testing.T) {

		req := &dto.ProfileUpdateRequest{Name: "Updated Name", PhoneNumber: "Updated Phone Number"}
		expectedUser := &repository.User{Id: 1, Name: "Updated Name", PhoneNumber: "Updated Phone Number"}
		user := &repository.User{}
		expectedErr := errors.New("repository error")

		mockUserRepo.On("FindByID", ctx, expectedUser.Id).Return(expectedUser, true, nil)
		mockUserRepo.On("FindByPhoneNumber", ctx, req.PhoneNumber).Return(user, false, expectedErr)
		mockUserRepo.On("Update", ctx, expectedUser).Return(expectedUser, nil)

		svc := service.NewUserService(mockUserRepo)
		UpdatedUser, err := svc.UpdateMyProfile(context.Background(), req, 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, UpdatedUser)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Failed to update profile", func(t *testing.T) {

		req := &dto.ProfileUpdateRequest{Name: "John Doe", PhoneNumber: "+6281209876543"}
		expectedUser := &repository.User{Id: 1, Name: "John Doe", PhoneNumber: "+6281234567890"}
		unexpectedUser := &repository.User{Id: 2, Name: "Borrys", PhoneNumber: "+6281209876543"}

		mockUserRepo.On("FindByID", ctx, expectedUser.Id).Return(expectedUser, true, nil)
		mockUserRepo.On("FindByPhoneNumber", ctx, req.PhoneNumber).Return(unexpectedUser, true, nil)

		svc := service.NewUserService(mockUserRepo)
		user, err := svc.UpdateMyProfile(context.Background(), req, 1)

		assert.Error(t, err)
		assert.NotEqual(t, expectedUser, user)
		mockUserRepo.AssertExpectations(t)
	})
}
