package handlers

import (
	"encoding/json"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
	"fmedeiros95/golang-rest-api/app/services"
	"fmedeiros95/golang-rest-api/app/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(db *core.Database) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

func (uh *UserHandler) AuthedUser(w http.ResponseWriter, r *http.Request) {
	user := utils.GetUserFromContext(r.Context())
	core.RespondWithJSON(w, http.StatusOK, user, "User found successfully")
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		core.RespondWithError(w, http.StatusInternalServerError, "Error while trying to list users")
		return
	}
	core.RespondWithJSON(w, http.StatusOK, users, "Users found successfully")
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err = uh.userService.CreateUser(&payload)
	if err != nil {
		core.RespondWithError(w, http.StatusInternalServerError, "Error while trying to create user")
		return
	}

	core.RespondWithJSON(w, http.StatusCreated, nil, "User created successfully")
}

func (uh *UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	user, err := uh.userService.FindUser(uint(id))
	if err != nil {
		core.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	core.RespondWithJSON(w, http.StatusOK, user, "User found successfully")
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err = uh.userService.UpdateUser(payload)
	if err != nil {
		core.RespondWithError(w, http.StatusInternalServerError, "Error while trying to update user")
		return
	}

	core.RespondWithJSON(w, http.StatusOK, nil, "User updated successfully")
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		core.RespondWithError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	user, err := uh.userService.FindUser(uint(id))
	if err != nil {
		core.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	err = uh.userService.DeleteUser(*user)
	if err != nil {
		core.RespondWithError(w, http.StatusInternalServerError, "Error while trying to delete user")
		return
	}

	core.RespondWithJSON(w, http.StatusOK, nil, "User deleted successfully")
}
