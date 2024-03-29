package domain

import (
	"context"
)

type SignupRequest struct {
	Email    string `json:"email,omitempty" jsonschema:"description=Email,format=email,title=Email,required"`
	Name     string `json:"name,omitempty" jsonschema:"description=Name,title=Name"`
	Password string `json:"password,omitempty" jsonschema:"description=Password,title=Password,required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
