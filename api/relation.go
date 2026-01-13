package api

import (
	"encoding/json"
	"net/http"
	"groupie-tracker/models"
)

type relationAPI struct {
	Index []models.Relation `json:"index"`
}

func GetRelations() ([]models.Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data relationAPI
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Index, nil
}

func GetRelationByID(id int) (models.Relation, error) {
	rels, err := GetRelations()
	if err != nil {
		return models.Relation{}, err
	}

	for _, r := range rels {
		if r.ID == id {
			return r, nil
		}
	}

	return models.Relation{}, nil
}
