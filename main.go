package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {

	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes précises
	mux.HandleFunc("/artists", handlers.Artists)
	mux.HandleFunc("/artist", handlers.Artist)
	mux.HandleFunc("/404", handlers.NotFound)

	// Route racine + 404
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.Home(w, r)
			return
		}
		handlers.NotFound(w, r)
	})

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
