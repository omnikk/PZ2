package utils

import (
	"encoding/json"
	"net/http"
)

// JSONError описывает структуру JSON-ошибки
type JSONError struct {
	Error string `json:"error"`
}

// WriteJSON записывает любой объект в ответ в формате JSON
func WriteJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// WriteErr записывает ошибку с кодом HTTP
func WriteErr(w http.ResponseWriter, code int, msg string) {
	WriteJSON(w, code, JSONError{Error: msg})
}