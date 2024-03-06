package handlers

import (
	"fmedeiros95/golang-rest-api/app/core"
	"net/http"
)

type AppMetadata struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	metadata := AppMetadata{
		Name:    "GoLang Rest API",
		Version: "0.0.1",
		Author:  "Felipe Medeiros <medeiros.dev@gmail.com>",
	}
	core.RespondWithJSON(w, http.StatusOK, metadata)
}
