package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"groupie-tracker/models"
)

// URL de l'API Groupie Tracker pour les relations artistes / concerts
const relationsURL = "https://groupietrackers.herokuapp.com/api/relation"

// GetRelationByArtistID retourne les lieux et dates de concerts d'un artiste
func GetRelationByArtistID(id int) (map[string][]string, error) {
	resp, err := http.Get(relationsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response models.RelationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	for _, relation := range response.Index {
		if relation.ID == id {
			return relation.DatesLocations, nil
		}
	}

	return nil, errors.New("no concert data found for this artist")
}
