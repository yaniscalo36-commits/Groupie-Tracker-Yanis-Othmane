package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/artists", handlers.Artists)
	http.HandleFunc("/artist", handlers.Artist)
	http.HandleFunc("/location", handlers.Location)




	log.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

