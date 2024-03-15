package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := Response{Message: message, StatusCode: statusCode}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	response := Response{Message: message, StatusCode: statusCode, Data: &data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
