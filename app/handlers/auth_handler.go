package handlers

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/services"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(db *core.Database) *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(db),
	}
}

func (ah *AuthHandler) AuthLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate user"))
}

func (ah *AuthHandler) AuthRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register a new user"))
}
