package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"homepage-mock-api/internal/mock"
	"homepage-mock-api/internal/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, mock.BuildHomeResponse())
}

func HomeComposed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// In a real app clientID comes from auth middleware.
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		clientID = "user_42"
	}

	home := services.NewHomeService()
	writeJSON(w, home.Build(clientID))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
