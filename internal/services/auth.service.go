package services

import (
	"crypto/sha1"
	"fmt"
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

const salt = "sdfsjlib81gbsdbfksdjf"

type AuthService struct {
	repo repository.Authorization
}

func InitAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// -- Helpers --

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// -- Functions --

func (service *AuthService) SignUpUser(user *model.UserSignUp) (uint, error) {

	user.Password = generatePasswordHash(user.Password)

	return service.repo.CreateUser(user)
}
