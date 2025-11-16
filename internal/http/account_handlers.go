package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/olliekm/realtime-ledger/internal/service"
)

func (h *Handlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req service.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	acct, err := h.ledgerSvc.CreateAccount(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(acct)
}

func (h *Handlers) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	acct, err := h.ledgerSvc.GetAccount(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(acct)
}

func (h *Handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bal, err := h.ledgerSvc.GetBalance(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(bal)
}

func (h *Handlers) GetEntries(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	filter, err := parseListFilter(r, id)
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

func parseListFilter(r *http.Request, accountID string) (service.ListEntriesFilter, error) {
	q := r.URL.Query()
	filter := service.ListEntriesFilter{AccountID: accountID}

	if from := q.Get("from"); from != "" {
		t, err := time.Parse(time.RFC3339, from)
		if err != nil {
			return filter, errors.New("invalid from param, expected RFC3339")
		}
		filter.From = &t
	}
	if to := q.Get("to"); to != "" {
		t, err := time.Parse(time.RFC3339, to)
		if err != nil {
			return filter, errors.New("invalid to param, expected RFC3339")
		}
		filter.To = &t
	}
	if limit := q.Get("limit"); limit != "" {
		v, err := strconv.Atoi(limit)
		if err != nil {
			return filter, errors.New("limit must be an integer")
		}
		filter.Limit = v
	}
	if offset := q.Get("offset"); offset != "" {
		v, err := strconv.Atoi(offset)
		if err != nil {
			return filter, errors.New("offset must be an integer")
		}
		filter.Offset = v
	}
	return filter, nil
}
