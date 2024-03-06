package repositories

import (
	"errors"
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/models"
)

type UserRepository struct {
	db *core.Database
}

func NewUserRepository(db *core.Database) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUserByID(userId uint) (*models.User, error) {
	var user models.User
	result := ur.db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	var user models.User
	result := ur.db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
