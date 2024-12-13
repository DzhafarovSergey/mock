package services

import (
	"errors"
	"fmt"
	"mock/config"
	"mock/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) ValidateCredentials(request models.OAuthRequest) (string, error) {
	conf := config.NewServiceConfig()
	validClientId := conf.ClientId
	validSecretKey := conf.SecretKey

	if request.ClientId != validClientId || request.SecretKey != validSecretKey {
		return "", errors.New("Invalid credentians")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"clientId": validClientId,
		"exp":      time.Now().Add(4 * time.Hour).Unix(),
	})
	fmt.Println("sadasdasfas", token)
	tokenString, err := token.SignedString([]byte(validSecretKey))

	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, err
}
