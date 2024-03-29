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

func TestGetUserByEmailLogin(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	email := "test@test.com"
	name := "John Doe"

	t.Run("success", func(t *testing.T) {
		mockUser := &domain.User{
			ID:    1,
			Name:  &name,
			Email: email,
		}

		mockUserRepository.On("FindByEmail", mock.Anything, email).Return(mockUser, nil).Once()

		u := NewLoginUsecase(mockUserRepository, time.Second*2)

		user, err := u.GetUserByEmail(context.Background(), email)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, email).Return(nil, errors.New("Unexpected")).Once()

		u := NewLoginUsecase(mockUserRepository, time.Second*2)

		user, err := u.GetUserByEmail(context.Background(), email)

		assert.Error(t, err)
		assert.Nil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestCreateAccessTokenLogin(t *testing.T) {
	name := "John Doe"
	user := &domain.User{
		ID:    1,
		Name:  &name,
		Email: "test@test.com",
	}

	secret := "secret"
	expiry := 3600

	u := NewLoginUsecase(nil, time.Second*2)

	token, err := u.CreateAccessToken(user, secret, expiry)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestCreateRefreshTokenLogin(t *testing.T) {
	name := "John Doe"
	user := &domain.User{
		ID:    1,
		Name:  &name,
		Email: "test@test.com",
	}

	secret := "secret"
	expiry := 3600

	u := NewLoginUsecase(nil, time.Second*2)

	token, err := u.CreateRefreshToken(user, secret, expiry)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewLoginUsecase(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	timeout := time.Second * 2

	u := NewLoginUsecase(mockUserRepository, timeout)

	assert.NotNil(t, u)
	assert.Equal(t, mockUserRepository, u.(*loginUsecase).userRepository)
	assert.Equal(t, timeout, u.(*loginUsecase).contextTimeout)
}
