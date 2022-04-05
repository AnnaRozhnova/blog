package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/AnnaRozhnova/blog/pkg/repository"
)

const (
	salt = "sdasldkjldfknvls"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user blog.User) error {

	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(username, password string) (blog.User, error) {
	passwordHash := generatePasswordHash(password)
	return s.repo.GetUser(username, passwordHash)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
