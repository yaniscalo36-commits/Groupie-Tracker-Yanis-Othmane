package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)


func matchMembers(count int, filters []string) bool {
	if len(filters) == 0 {
		return true
	}

	for _, f := range filters {
		switch f {
		case "1":
			if count == 1 {
				return true
			}
		case "2":
			if count == 2 {
				return true
			}
		case "3":
			if count >= 3 && count <= 4 {
				return true
			}
		case "5":
			if count >= 5 {
				return true
			}
		}
	}
	return false
}

func Artists(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("search"))
	minYearStr := r.URL.Query().Get("minYear")
	maxYearStr := r.URL.Query().Get("maxYear")
	membersFilters := r.URL.Query()["members"]

	minYear := 0
	maxYear := 3000

	if minYearStr != "" {
		minYear, _ = strconv.Atoi(minYearStr)
	}
	if maxYearStr != "" {
		maxYear, _ = strconv.Atoi(maxYearStr)
	}

	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Erreur API", 500)
		return
	}

	var filtered []models.Artist

	for _, a := range artists {

		// filtre année
		if a.CreationDate < minYear || a.CreationDate > maxYear {
			continue
		}

		// filtre nombre de membres
		if !matchMembers(len(a.Members), membersFilters) {
			continue
		}

		// filtre recherche
		if query != "" {
			if strings.Contains(strings.ToLower(a.Name), query) {
				filtered = append(filtered, a)
				continue
			}

			for _, m := range a.Members {
				if strings.Contains(strings.ToLower(m), query) {
					filtered = append(filtered, a)
					break
				}
			}
		} else {
			filtered = append(filtered, a)
		}
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	tmpl.Execute(w, filtered)
}
