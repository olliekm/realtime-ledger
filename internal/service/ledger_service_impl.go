package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/olliekm/realtime-ledger/internal/ledger"
)

type ledgerService struct {
	store LedgerStore
}

func NewLedgerService(store LedgerStore) LedgerService {
	return &ledgerService{
		store: store,
	}
}

func (s *ledgerService) CreateAccount(ctx context.Context, req CreateAccountRequest) (*ledger.Account, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}
	currency, err := parseCurrency(req.Currency)
	if err != nil {
		return nil, err
	}

	acct := &ledger.Account{
		Name:      req.Name,
		Currency:  currency,
		CreatedAt: time.Now(),
	}

	if err := s.store.CreateAccount(ctx, acct); err != nil {
		return nil, err
	}
	return acct, nil
}

func (s *ledgerService) GetAccount(ctx context.Context, id string) (*ledger.Account, error) {
	return s.store.GetAccount(ctx, id)
}

func (s *ledgerService) GetBalance(ctx context.Context, id string) (ledger.Money, error) {
	return s.store.GetBalance(ctx, id)
}

func (s *ledgerService) PostJournal(ctx context.Context, req PostJournalRequest) (*ledger.Entry, error) {
	if len(req.Entries) == 0 {
		return nil, errors.New("journal must contain at least one entry")
	}

	entry := ledger.Entry{
		EffectiveAt: req.At,
		CreatedAt:   time.Now(),
	}

	for _, e := range req.Entries {
		account, err := s.store.GetAccount(ctx, e.AccountID)
		if err != nil {
			return nil, fmt.Errorf("account %s: %w", e.AccountID, err)
		}
		amount := e.Amount
		if strings.EqualFold(e.Side, "credit") {
			amount = -amount
		}
		entry.Postings = append(entry.Postings, ledger.Posting{
			AccountID: ledger.AccountID(e.AccountID),
			Amount:    ledger.NewMoney(amount, account.Currency),
		})
	}

	journal := &ledger.Journal{
		Entries:     []ledger.Entry{entry},
		EffectiveAt: req.At,
		CreatedAt:   time.Now(),
		Metadata:    map[string]string{},
	}
	if err := s.store.InsertJournal(ctx, journal); err != nil {
		return nil, err
	}

	// return the first entry for now until batch support is needed
	return &journal.Entries[0], nil
}

func (s *ledgerService) ListEntries(ctx context.Context, filter ListEntriesFilter) ([]*ledger.Entry, error) {
	entries, err := s.store.ListEntries(ctx, filter.AccountID, filter.From, filter.To, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*ledger.Entry, 0, len(entries))
	for i := range entries {
		entry := entries[i]
		result = append(result, &entry)
	}
	return result, nil
}

func parseCurrency(code string) (ledger.Currency, error) {
	switch strings.ToUpper(code) {
	case string(ledger.USD):
		return ledger.USD, nil
	case string(ledger.EUR):
		return ledger.EUR, nil
	case string(ledger.GBP):
		return ledger.GBP, nil
	default:
		return "", fmt.Errorf("unsupported currency %q", code)
	}
}
