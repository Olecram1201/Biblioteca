package controllers

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {

	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		return err
	}

	// Render template
	t.Execute(w, data)

	return nil
}
