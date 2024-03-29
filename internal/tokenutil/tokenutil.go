package tokenutil

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jordanlanch/ionix-test/domain"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.JwtCustomClaims{
		Email: user.Email,
		ID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(tokenString, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Cambia "userID" por "id" para que coincida con cómo guardas el ID en el token
		userID, ok := claims["id"]
		if !ok {
			return 0, fmt.Errorf("id no encontrado en el token")
		}

		// Manejar la conversión desde float64 a int
		switch v := userID.(type) {
		case float64:
			return int(v), nil
		case int:
			return v, nil
		default:
			return 0, fmt.Errorf("tipo de id no esperado: %T", v)
		}
	}

	return 0, fmt.Errorf("token inválido o expirado")
}
