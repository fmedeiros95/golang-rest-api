package middlewares

import (
	"errors"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
	"fmedeiros95/golang-rest-api/app/services"
	"fmedeiros95/golang-rest-api/app/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	userService *services.UserService
}

func NewAuthMiddleware(db *core.Database) *AuthMiddleware {
	return &AuthMiddleware{
		userService: services.NewUserService(db),
	}
}

func (am *AuthMiddleware) ValidateToken(tokenString string) (*models.User, error) {
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

	userID, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("invalid token")
	}

	// ** Obtendo o usuário do banco de dados
	user, err := am.userService.FindUser(uint(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (am *AuthMiddleware) JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ** Get token from headers
		authorization := strings.Split(r.Header.Get("Authorization"), " ")
		if authorization[0] != "Bearer" {
			core.RespondWithError(w, http.StatusUnauthorized, "Invalid token type")
			return
		}

		// ** Check if the header has value
		if authorization[1] == "" {
			core.RespondWithError(w, http.StatusUnauthorized, "Authorization token is missing")
			return
		}

		// ** Check if token is valid
		user, err := am.ValidateToken(authorization[1])
		if err != nil {
			core.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// ** Attach user to context
		ctx := utils.SetUserInContext(r.Context(), user)

		// ** Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
