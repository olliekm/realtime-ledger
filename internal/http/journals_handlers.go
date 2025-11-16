package api

import (
	"encoding/json"
	"net/http"

	"github.com/olliekm/realtime-ledger/internal/service"
)

func (h *Handlers) PostJournal(w http.ResponseWriter, r *http.Request) {
	var req service.PostJournalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	entry, err := h.ledgerSvc.PostJournal(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(entry)
}

func (h *Handlers) ListEntries(w http.ResponseWriter, r *http.Request) {
	filter, err := parseListFilter(r, r.URL.Query().Get("account_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entries, err := h.ledgerSvc.ListEntries(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(entries)
}
