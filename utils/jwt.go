package utils

import (
	"errors"
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/models"
	"fmedeiros95/golang-rest-api/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(core.Config.JwtSecret))
}

func ValidateToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(core.Config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return nil, errors.New("invalid token")
	}

	userRepository := repositories.NewUserRepository(&core.Database{})
	user, err := userRepository.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
