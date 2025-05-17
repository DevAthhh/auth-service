package services

import (
	"time"

	"github.com/DevAthhh/auth-service/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	secretKey []byte
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{
		secretKey: []byte(secretKey),
	}
}

type customClaims struct {
	UserID   uuid.UUID
	Username string
	Email    string
	jwt.RegisteredClaims
}

func (a *AuthService) GenerateToken(user *models.User) (string, error) {
	claims := customClaims{
		UserID:   user.GetID(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 15 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secretKey))
}

func (a *AuthService) ValidateToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*customClaims), nil

}
