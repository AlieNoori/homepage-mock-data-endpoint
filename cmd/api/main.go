package main

import (
	"log"
	"net/http"
	"time"

	"homepage-mock-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/home", handlers.Home)
	mux.HandleFunc("/api/v1/home/composed", handlers.HomeComposed)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("mock home API on localhost:8080")
	log.Println("  GET /api/v1/home")
	log.Println("  GET /api/v1/home/composed")
	log.Fatal(srv.ListenAndServe())
}
