package api

import (
	"net/http"

	"github.com/olliekm/realtime-ledger/internal/service"
)

type Handlers struct {
	ledgerSvc service.LedgerService
}

func NewHandlers(ledgerSvc service.LedgerService) *Handlers {
	return &Handlers{ledgerSvc: ledgerSvc}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
