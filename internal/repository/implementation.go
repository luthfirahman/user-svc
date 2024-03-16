package repository

import (
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

// Save implements IUserRepository.
func (r *userRepository) Save(ctx context.Context, user *User) (*User, error) {

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return &User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*User, bool, error) {
	var user *User
	if err := r.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).Take(&user).Error; err != nil {
		return &User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*User, bool, error) {
	var user *User
	if err := r.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error; err != nil {
		return &User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) Update(ctx context.Context, user *User) (*User, error) {
	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		return &User{}, err
	}

	return user, nil
}
