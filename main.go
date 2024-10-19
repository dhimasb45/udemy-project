package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)

	// Server
	log.Printf("Starting application on port %s", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
