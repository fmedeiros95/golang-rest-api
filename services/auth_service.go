package services

import (
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/repositories"
	"fmedeiros95/golang-rest-api/utils"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository: repositories.NewUserRepository(&core.Database{}),
	}
}

func (as *AuthService) Login(username, password string) (string, error) {
	user, err := as.userRepository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	return utils.CreateToken(user.ID)
}
