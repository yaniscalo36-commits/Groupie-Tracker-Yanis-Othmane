package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}
