package handlers

import (
	"fmedeiros95/golang-rest-api/app/core"
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
	core.RespondWithJSON(w, http.StatusOK, user)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// ** Obtém os parâmetros de consulta para paginação, se fornecidos
	pageNumStr := r.URL.Query().Get("page")
	pageLimitStr := r.URL.Query().Get("limit")

	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		pageNum = 1
	}
	pageLimit, err := strconv.Atoi(pageLimitStr)
	if err != nil {
		pageLimit = 10
	}

	users, err := uh.userService.ListUsers(pageNum, pageLimit)
	if err != nil {
		core.RespondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}
	core.RespondWithJSON(w, http.StatusOK, users)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
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

	core.RespondWithJSON(w, http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update unique user"))
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete unique user"))
}
