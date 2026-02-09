package models
// Cette structure est utilisée pour afficher les informations
// sur la page des artistes et la page de détail.
type Artist struct {
	ID           int      `json:"id"`           // Identifiant unique de l'artiste
	Name         string   `json:"name"`         // Nom de l'artiste ou du groupe
	Image        string   `json:"image"`        // Image officielle
	CreationDate int      `json:"creationDate"` // Année de création
	FirstAlbum   string   `json:"firstAlbum"`   // Date du premier album
	Members      []string `json:"members"`      // Liste des membres du groupe
}
