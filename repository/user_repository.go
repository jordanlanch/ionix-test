package repository

import (
	"context"

	"github.com/jordanlanch/ionix-test/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db    *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	result := r.db.WithContext(ctx).First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	result := r.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
