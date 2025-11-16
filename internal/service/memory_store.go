package service

import (
	"context"
	"time"

	"github.com/olliekm/realtime-ledger/internal/ledger"
)

// InMemoryStore provides a simple LedgerStore backed by the in-memory ledger.
type InMemoryStore struct {
	ledger *ledger.Ledger
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{ledger: ledger.NewInMemoryLedger()}
}

func (s *InMemoryStore) CreateAccount(ctx context.Context, a *ledger.Account) error {
	acct, err := s.ledger.CreateAccount(a.Name, a.Currency)
	if err != nil {
		return err
	}
	*a = *acct
	return nil
}

func (s *InMemoryStore) GetAccount(ctx context.Context, id string) (*ledger.Account, error) {
	return s.ledger.GetAccount(ledger.AccountID(id))
}

func (s *InMemoryStore) InsertJournal(ctx context.Context, j *ledger.Journal) error {
	for i := range j.Entries {
		entry := j.Entries[i]
		if _, err := s.ledger.PostEntry(ctx, &entry); err != nil {
			return err
		}
	}
	return nil
}

func (s *InMemoryStore) ListEntries(ctx context.Context, accountID string, from, to *time.Time, limit, offset int) ([]ledger.Entry, error) {
	entries := s.ledger.ListEntries(ledger.AccountID(accountID))

	// Basic filtering by time window.
	filtered := make([]ledger.Entry, 0, len(entries))
	for _, e := range entries {
		if from != nil && e.EffectiveAt.Before(*from) {
			continue
		}
		if to != nil && e.EffectiveAt.After(*to) {
			continue
		}
		filtered = append(filtered, e)
	}

	// Apply offset/limit slicing.
	start := offset
	if start > len(filtered) {
		return []ledger.Entry{}, nil
	}
	end := len(filtered)
	if limit > 0 && start+limit < end {
		end = start + limit
	}
	return filtered[start:end], nil
}

func (s *InMemoryStore) GetBalance(ctx context.Context, accountID string) (ledger.Money, error) {
	return s.ledger.GetBalance(ledger.AccountID(accountID))
}
