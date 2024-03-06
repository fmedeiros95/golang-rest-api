package handlers

import (
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/models"
	"net/http"
)

func UserAuthed(w http.ResponseWriter, r *http.Request) {
	core.RespondWithJSON(w, http.StatusOK, models.User{FirstName: "Felipe", LastName: "Medeiros", Email: "medeiros.dev@gmail.com"})
}

func UsersList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list"))
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new user"))
}

func UsersFind(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find unique user"))
}

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update unique user"))
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete unique user"))
}
