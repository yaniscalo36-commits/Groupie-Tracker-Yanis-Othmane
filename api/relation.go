package api

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/models"
)

const relationsURL = "https://groupietrackers.herokuapp.com/api/relation"

type relationResponse struct {
	Index []models.Relation `json:"index"`
}

func GetRelations() ([]models.Relation, error) {
	resp, err := http.Get(relationsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res relationResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res.Index, nil
}

func GetRelationByID(id int) (models.Relation, error) {
	list, err := GetRelations()
	if err != nil {
		return models.Relation{}, err
	}

	for _, r := range list {
		if r.ID == id {
			return r, nil
		}
	}
	return models.Relation{}, nil
}
