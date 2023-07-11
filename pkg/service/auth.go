package service

import (
	"crypto/sha1"
	"fmt"
	restapi "rest_api_TODO"
	"rest_api_TODO/pkg/repository"
)

const salt = "adksjdfl234skjd"

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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
