package utils

import (
	"fmedeiros95/golang-rest-api/app/core"
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
