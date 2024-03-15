package validations

import (
	"errors"
	"net/mail"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthLoginValidation(data Login) (*Login, error) {
	// ** Validate empty fields
	if data.Email == "" || data.Password == "" {
		return nil, errors.New("fill in all fields")
	}

	// ** Validate email
	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	return &data, nil
}

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func AuthRegisterValidation(data Register) (*Register, error) {
	if data.FirstName == "" || data.LastName == "" || data.Email == "" || data.Password == "" {
		return nil, errors.New("fill in all fields")
	}

	// ** Validate name
	if len(data.FirstName) < 3 || len(data.LastName) < 3 {
		return nil, errors.New("name and last name must contain at least 3 characters")
	}

	// ** Validate email
	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	// ** Validate password
	if len(data.Password) < 8 {
		return nil, errors.New("password must contain at least 8 characters")
	}

	return &data, nil
}
