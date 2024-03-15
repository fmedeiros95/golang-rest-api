package services

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
	"fmedeiros95/golang-rest-api/app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(db *core.Database) *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (us *UserService) ListUsers() (*[]models.User, error) {
	users, err := us.userRepository.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) FindUser(userId uint) (*models.User, error) {
	user, err := us.userRepository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.userRepository.CreateUser(user)
}

func (us *UserService) UpdateUser(user models.User) error {
	return us.userRepository.UpdateUser(user)
}

func (us *UserService) DeleteUser(user models.User) error {
	return us.userRepository.DeleteUser(user)
}
