package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/models"
)

const artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func GetArtistByID(id int) (models.Artist, error) {
	list, err := GetArtists()
	if err != nil {
		return models.Artist{}, err
	}

	for _, a := range list {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Artist{}, fmt.Errorf("artist not found")
}
