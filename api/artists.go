package api

import (
	"encoding/json"
	"net/http"
	"groupie-tracker/models"
	"fmt"
)

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
func GetArtistByID(id int) (models.Artist, error) {
	artists, err := GetArtists()
	if err != nil {
		return models.Artist{}, err
	}

	for _, a := range artists {
		if a.ID == id {
			return a, nil
		}
	}

	return models.Artist{}, fmt.Errorf("not found")
}
