package handlers

import (
	"groupie-tracker/api"
	"html/template"
	"net/http"
)

func Location(w http.ResponseWriter, r *http.Request) {
	place := r.URL.Query().Get("place")
	if place == "" {
		http.Error(w, "Lieu manquant", 400)
		return
	}

	relations, err := api.GetRelations()
	if err != nil {
		http.Error(w, "Erreur API", 500)
		return
	}

	type Concert struct {
		ArtistID int
		Dates    []string
	}

	var concerts []Concert

	for _, rel := range relations {
		if rel.DatesLocations != nil {
			if dates, ok := rel.DatesLocations[place]; ok {
				concerts = append(concerts, Concert{
					ArtistID: rel.ID,
					Dates:    dates,
				})
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/location.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	tmpl.Execute(w, concerts)
}
