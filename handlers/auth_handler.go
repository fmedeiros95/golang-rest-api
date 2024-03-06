package handlers

import "net/http"

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticate user"))
}

func AuthRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register a new user"))
}

func AuthGetAuthedUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get authed user info"))
}
