package models

// Relation représente les concerts associés à un artiste.
// Chaque lieu est lié à une liste de dates.
type Relation struct {
	ID             int                 `json:"id"`             // Identifiant de l'artiste
	DatesLocations map[string][]string `json:"datesLocations"` // Lieu -> dates de concert
}

type RelationsResponse struct {
	Index []Relation `json:"index"`
}
