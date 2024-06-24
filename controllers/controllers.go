package controllers

import (
	"html/template"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error.html", nil)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "not-found.html", nil)
}
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Método no permitido"))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {

	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		return err
	}

	// Render template
	t.Execute(w, data)

	return nil
}

// EXCEPTIONS
