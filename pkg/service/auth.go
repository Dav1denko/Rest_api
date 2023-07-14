package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "adksjdfl234skjd"
	tokenTTL   = 12 * time.Hour
	signingKey = "adsfsfdwdf342sdfsdfs123asdfSSDFdfsdfsdf"
)

type tokenClains struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authoruzation
}

func NewAuthService(repo repository.Authoruzation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user restapi.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClains{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClains{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClains)
	if !ok {
		return 0, errors.New("token clains are not of type tokenClains")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
