package controllers

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "error.html", nil)
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "not-found.html", nil)
}
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Método no permitido"))
}

// EXCEPTIONS
