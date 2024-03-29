package domain

import (
	"context"
)

type LoginRequest struct {
	Email    string `json:"email,omitempty" jsonschema:"required,description=Email of the user,title=Email"`
	Password string `json:"password,omitempty" jsonschema:"required,description=Password of the user,title=Password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
