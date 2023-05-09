package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/Marif226/go-todo-rest/internal/repository"
)

const salt = "kbL3FqE6mbMx6xhwnZYpCjKmE8rHl2Qy"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	// replace password with generated hash
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}