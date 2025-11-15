package main

import (
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/olliekm/realtime-ledger/internal/http/handlers"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	amw := Auth{tokenUsers: make(map[string]string)}
	amw.Populate()

	r.Use(amw.Middleware)

	r.HandleFunc("/v1/accounts", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/v1/accounts/{id}", handlers.GetAccount).Methods("GET")
	r.HandleFunc("/v1/accounts/{id}/balance", handlers.GetBalance).Methods("GET")
	r.HandleFunc("/v1/accounts/{id}/entires", handlers.GetEntries).Methods("GET")

	r.HandleFunc("/v1/journals", handle.PostJournal).Methods("POST")

	return r

}
