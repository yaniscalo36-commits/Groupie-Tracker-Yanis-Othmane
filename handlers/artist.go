package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", 400)
		return
	}

	artist, err := api.GetArtistByID(id)
	if err != nil {
		http.Error(w, "Artiste introuvable", 404)
		return
	}

	relation, err := api.GetRelationByID(id)
	if err != nil {
		http.Error(w, "Concerts introuvables", 500)
		return
	}

	// On regroupe tout dans une seule structure
	data := struct {
		Artist   models.Artist
		Relation models.Relation
	}{
		Artist:   artist,
		Relation: relation,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	tmpl.Execute(w, data)
}
