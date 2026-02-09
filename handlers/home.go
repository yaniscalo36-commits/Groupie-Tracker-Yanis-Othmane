package handlers

import (
	"html/template"
	"net/http"
)

// Home affiche la page d'accueil du site
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page d'accueil", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
