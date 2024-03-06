package services

import (
	"errors"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/repositories"
	"fmedeiros95/golang-rest-api/app/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(db *core.Database) *AuthService {
	return &AuthService{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		return "", errors.New("password missmatch")
	}

	return utils.CreateToken(user.ID)
}
