package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {

	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artists", handlers.Artists)
	http.HandleFunc("/artist", handlers.Artist)
	http.HandleFunc("/location", handlers.Location)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
