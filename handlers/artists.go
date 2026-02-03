package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

func matchMembers(n int, filters []string) bool {
	if len(filters) == 0 {
		return true
	}
	for _, f := range filters {
		switch f {
		case "1":
			if n == 1 {
				return true
			}
		case "2":
			if n == 2 {
				return true
			}
		case "3":
			if n >= 3 && n <= 4 {
				return true
			}
		case "5":
			if n >= 5 {
				return true
			}
		}
	}
	return false
}

func Artists(w http.ResponseWriter, r *http.Request) {
	q := strings.ToLower(r.URL.Query().Get("search"))
	minY, _ := strconv.Atoi(r.URL.Query().Get("minYear"))
	maxY, _ := strconv.Atoi(r.URL.Query().Get("maxYear"))
	if maxY == 0 {
		maxY = 3000
	}
	fMembers := r.URL.Query()["members"]

	list, err := api.GetArtists()
	if err != nil {
		http.Error(w, "api error", 500)
		return
	}

	var res []models.Artist

	for _, a := range list {
		if a.CreationDate < minY || a.CreationDate > maxY {
			continue
		}
		if !matchMembers(len(a.Members), fMembers) {
			continue
		}
		if q != "" {
			ok := strings.Contains(strings.ToLower(a.Name), q)
			for _, m := range a.Members {
				if strings.Contains(strings.ToLower(m), q) {
					ok = true
				}
			}
			if !ok {
				continue
			}
		}
		res = append(res, a)
	}

	tpl, _ := template.ParseFiles("templates/artists.html")
	tpl.Execute(w, res)
}
