package api

import (
	"encoding/json"
	"net/http"
)

type CreateAccountRequest struct {
	Name string `json:"name"`
}

func (h *Handlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating an account
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
}
