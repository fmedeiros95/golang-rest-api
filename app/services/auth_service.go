package services

import (
	"errors"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
	"fmedeiros95/golang-rest-api/app/repositories"
	"fmedeiros95/golang-rest-api/app/utils"
	"fmedeiros95/golang-rest-api/app/validations"

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

type AuthResponse struct {
	Token *string `json:"token,omitempty"`
}

func (as *AuthService) Login(payload validations.Login) (*AuthResponse, error) {
	data, err := validations.AuthLoginValidation(payload)
	if err != nil {
		return nil, err
	}

	user, err := as.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return nil, errors.New("email and/or password invalid")
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)) != nil {
		return nil, errors.New("email and/or password invalid")
	}

	jwtToken, err := utils.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{Token: &jwtToken}, nil
}

func (as *AuthService) Register(payload validations.Register) (*AuthResponse, error) {
	data, err := validations.AuthRegisterValidation(payload)
	if err != nil {
		return nil, err
	}

	// ** Verificar se o email já está cadastrado
	_, err = as.userRepository.GetUserByEmail(data.Email)
	if err == nil {
		return nil, errors.New("email already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashedPassword,
	}
	err = as.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{}, nil
}
