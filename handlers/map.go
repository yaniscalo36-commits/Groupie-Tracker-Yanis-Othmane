package handlers

import (
	"groupie-tracker/api"
	"html/template"
	"net/http"
)

func Map(w http.ResponseWriter, r *http.Request) {
	relations, err := api.GetRelations()
	if err != nil {
		http.Error(w, "Erreur API", 500)
		return
	}

	tmpl, _ := template.ParseFiles("templates/map.html")
	tmpl.Execute(w, relations)
}
