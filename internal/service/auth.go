package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Marif226/go-todo-rest/internal/model"
	"github.com/Marif226/go-todo-rest/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt = "kbL3FqE6mbMx6xhwnZYpCjKmE8rHl2Qy"
	secretKey = "N40Xzo-2DRUx_3xBgEJQDFM9Utn5f7tyf9lvFaPcy9Y"
)

type tokenClaims struct {
	jwt.Claims
	UserId	int	`json:"user_id"`
}

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
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.MapClaims{
			"exp" : time.Now().Add(12 * time.Hour).Unix(),
			"iat" : time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(secretKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
