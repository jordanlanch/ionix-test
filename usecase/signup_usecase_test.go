package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/jordanlanch/ionix-test/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	name := "John Doe"
	newUser := &domain.User{
		ID:    1,
		Name:  &name,
		Email: "test@test.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Create", mock.Anything, newUser).Return(newUser, nil).Once()

		u := NewSignupUsecase(mockUserRepository, time.Second*2)

		user, err := u.Create(context.Background(), newUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, newUser.Email, user.Email)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("Create", mock.Anything, newUser).Return(nil, errors.New("Unexpected")).Once()

		u := NewSignupUsecase(mockUserRepository, time.Second*2)

		user, err := u.Create(context.Background(), newUser)

		assert.Error(t, err)
		assert.Nil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetUserByEmail(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	email := "test@test.com"

	t.Run("success", func(t *testing.T) {
		name := "John Doe"
		mockUser := &domain.User{
			ID:    1,
			Name:  &name,
			Email: email,
		}

		mockUserRepository.On("FindByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		u := NewSignupUsecase(mockUserRepository, time.Second*2)

		user, err := u.GetUserByEmail(context.Background(), email)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, email).Return(nil, errors.New("Unexpected")).Once()

		u := NewSignupUsecase(mockUserRepository, time.Second*2)

		user, err := u.GetUserByEmail(context.Background(), email)

		assert.Error(t, err)
		assert.Nil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestCreateAccessToken(t *testing.T) {
	name := "John Doe"
	user := &domain.User{
		ID:    1,
		Name:  &name,
		Email: "test@test.com",
	}

	secret := "secret"
	expiry := 3600

	u := NewSignupUsecase(nil, time.Second*2)

	token, err := u.CreateAccessToken(user, secret, expiry)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestCreateRefreshToken(t *testing.T) {
	name := "John Doe"
	user := &domain.User{
		ID:    1,
		Name:  &name,
		Email: "test@test.com",
	}

	secret := "secret"
	expiry := 3600

	u := NewSignupUsecase(nil, time.Second*2)

	token, err := u.CreateRefreshToken(user, secret, expiry)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewSignupUsecase(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	timeout := time.Second * 2

	u := NewSignupUsecase(mockUserRepository, timeout)

	assert.NotNil(t, u)
	assert.Equal(t, mockUserRepository, u.(*signupUsecase).userRepository)
	assert.Equal(t, timeout, u.(*signupUsecase).contextTimeout)
}
