package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

// ConcertPoint représente un concert affiché dans la liste
// (un lieu + une date)
type ConcertPoint struct {
	Place string
	Date  string
}

// Artist affiche la page de détail d'un artiste
func Artist(w http.ResponseWriter, r *http.Request) {

	// --- Récupération et validation de l'ID ---
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Identifiant d'artiste invalide", http.StatusBadRequest)
		return
	}

	// --- Données de l'artiste ---
	artist, err := api.GetArtistByID(id)
	if err != nil {
		http.Error(w, "Artiste introuvable", http.StatusNotFound)
		return
	}

	// --- Données des concerts ---
	relations, err := api.GetRelationByArtistID(id)
	if err != nil {
		http.Error(w, "Aucune information de concert disponible", http.StatusInternalServerError)
		return
	}

	// Transformation des données pour l'affichage
	var concerts []ConcertPoint
	for place, dates := range relations {
		for _, date := range dates {
			concerts = append(concerts, ConcertPoint{
				Place: place,
				Date:  date,
			})
		}
	}

	// Données envoyées au template
	pageData := struct {
		Artist models.Artist
		Points []ConcertPoint
	}{
		Artist: artist,
		Points: concerts,
	}

	// Affichage du template
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page artiste", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, pageData)
}
