package repositories

import (
	"errors"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
)

type UserRepository struct {
	db *core.Database
}

func NewUserRepository(db *core.Database) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) ListUsers(pageNum, pageLimit int) (*[]models.User, error) {
	var users []models.User
	offset := (pageNum - 1) * pageLimit
	result := ur.db.Offset(offset).Limit(pageLimit).Find(&users)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &users, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUserByID(userId uint) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, userId)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(username string) (*models.User, error) {
	var user models.User
	result := ur.db.Where("email = ?", username).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (ur *UserRepository) DeleteUser(user models.User) error {
	return ur.db.Delete(user).Error
}
