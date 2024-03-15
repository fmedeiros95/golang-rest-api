package handlers

import (
	"encoding/json"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/services"
	"fmedeiros95/golang-rest-api/app/validations"
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
	var payload validations.Login
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	data, err := ah.authService.Login(payload)
	if err != nil {
		core.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	core.RespondWithJSON(w, http.StatusOK, data, "Login has been successful")
}

func (ah *AuthHandler) AuthRegister(w http.ResponseWriter, r *http.Request) {
	var payload validations.Register
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err = ah.authService.Register(payload)
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	core.RespondWithJSON(w, http.StatusCreated, nil, "User has been created")
}
