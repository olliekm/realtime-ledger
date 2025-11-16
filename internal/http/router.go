package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *Handlers) http.Handler {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	amw := &Auth{tokenUsers: make(map[string]string)}
	amw.Populate()

	api.Use(amw.Middleware)

	api.HandleFunc("/accounts", h.CreateAccount).Methods("POST")
	api.HandleFunc("/accounts/{id}", h.GetAccount).Methods("GET")
	api.HandleFunc("/accounts/{id}/balance", h.GetBalance).Methods("GET")
	api.HandleFunc("/accounts/{id}/entries", h.GetEntries).Methods("GET")

	api.HandleFunc("/journals", h.PostJournal).Methods("POST") // post a journal batch
	api.HandleFunc("/entries", h.ListEntries).Methods("GET")
	r.HandleFunc("/healthz", h.HealthCheck).Methods("GET")

	return r

}
