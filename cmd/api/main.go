package main

import (
	"log"
	"net/http"
	"time"

	"homepage-mock-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/home", handler.Home)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("mock home API listening on :8080 — try GET /api/v1/home")
	log.Fatal(srv.ListenAndServe())
}
