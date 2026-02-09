package handlers

import (
	"html/template"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("templates/404.html"))
	tmpl.Execute(w, nil)
}
