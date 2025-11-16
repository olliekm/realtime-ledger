package main

import (
	"net/http"

	router "github.com/olliekm/realtime-ledger/internal/http/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8080", r)

}
