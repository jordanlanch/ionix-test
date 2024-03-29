package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims representa los claims personalizados que se incluyen en el JWT de la plataforma .
type JwtCustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}

// JwtCustomRefreshClaims representa los claims personalizados que se incluyen en el JWT de actualización en la plataforma .
type JwtCustomRefreshClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
