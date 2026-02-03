package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/api"
)

type Concert struct {
	ArtistID int
	Dates    []string
}

func Location(w http.ResponseWriter, r *http.Request) {
	place := r.URL.Query().Get("place")
	if place == "" {
		http.Error(w, "missing place", http.StatusBadRequest)
		return
	}

	rels, err := api.GetRelations()
	if err != nil {
		http.Error(w, "api error", 500)
		return
	}

	var list []Concert

	for _, r := range rels {
		if dates, ok := r.DatesLocations[place]; ok {
			list = append(list, Concert{r.ID, dates})
		}
	}

	tpl, _ := template.ParseFiles("templates/location.html")
	tpl.Execute(w, list)
}
