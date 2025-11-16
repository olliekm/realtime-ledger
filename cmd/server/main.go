package main

import (
	"log"
	"net/http"

	api "github.com/olliekm/realtime-ledger/internal/http"
	"github.com/olliekm/realtime-ledger/internal/service"
)

func main() {
	store := service.NewInMemoryStore()
	svc := service.NewLedgerService(store)
	handlers := api.NewHandlers(svc)

	r := api.NewRouter(handlers)
	log.Fatal(http.ListenAndServe(":6767", r))

}
