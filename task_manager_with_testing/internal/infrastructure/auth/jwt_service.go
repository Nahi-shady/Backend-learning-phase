package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ParseToken(tokenString string) (string, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{secretKey: secretKey}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if ok && token.Valid {
		return claims.UserID, nil
	}
	return "", errors.New("invalid token")
}
