package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

// Artists affiche la page des artistes avec les filtres appliqués
func Artists(w http.ResponseWriter, r *http.Request) {

	//  Recherche texte (nom de l'artiste ou membre)
	search := strings.ToLower(r.URL.Query().Get("search"))

	//  Filtres par année de création
	minYear := 0
	maxYear := 3000

	if value := r.URL.Query().Get("minYear"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			minYear = year
		}
	}

	if value := r.URL.Query().Get("maxYear"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			maxYear = year
		}
	}

	//  Filtre par nombre de membres
	selectedMembers := r.URL.Query()["members"]

	// Récupération des artistes depuis l'API
	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Impossible de récupérer les artistes", http.StatusInternalServerError)
		return
	}

	var filtered []models.Artist

	for _, artist := range artists {

		// --- Filtre année ---
		if artist.CreationDate < minYear || artist.CreationDate > maxYear {
			continue
		}

		// --- Filtre membres ---
		if len(selectedMembers) > 0 {
			match := false

			for _, m := range selectedMembers {
				switch m {
				case "1":
					match = len(artist.Members) == 1
				case "2":
					match = len(artist.Members) == 2
				case "3":
					match = len(artist.Members) >= 3 && len(artist.Members) <= 4
				case "5":
					match = len(artist.Members) >= 5
				}
				if match {
					break
				}
			}

			if !match {
				continue
			}
		}

		// --- Recherche texte ---
		if search != "" {
			found := strings.Contains(strings.ToLower(artist.Name), search)

			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), search) {
					found = true
					break
				}
			}

			if !found {
				continue
			}
		}

		filtered = append(filtered, artist)
	}

	// Affichage du template
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page artistes", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, filtered)
}
