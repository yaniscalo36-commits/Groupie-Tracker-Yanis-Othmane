package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"groupie-tracker/models"
)

// URL de l'API Groupie Tracker pour les artistes
const artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

// GetArtists récupère la liste complète des artistes depuis l'API
func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// GetArtistByID retourne un artiste précis à partir de son ID
// (en parcourant la liste déjà fournie par l'API)
func GetArtistByID(id int) (models.Artist, error) {
	artists, err := GetArtists()
	if err != nil {
		return models.Artist{}, err
	}

	for _, artist := range artists {
		if artist.ID == id {
			return artist, nil
		}
	}

	return models.Artist{}, errors.New("artist not found")
}
