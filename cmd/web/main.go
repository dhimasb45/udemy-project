package main

import (
	"log"
	"net/http"

	"github.com/dhimasb45/udemy-project/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Server
	log.Printf("Starting application on port %s", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
