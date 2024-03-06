package services

import (
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/models"
	"fmedeiros95/golang-rest-api/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(&core.Database{}),
	}
}

func (us *UserService) Register(user *models.User) error {
	return us.userRepository.CreateUser(user)
}
