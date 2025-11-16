package ledger

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Ledger struct {
	// Implementation details would go here.
	mu       sync.RWMutex
	accounts map[AccountID]*Account
	balances map[AccountID]Money
	entries  map[EntryID]*Entry
}

func NewInMemoryLedger() *Ledger {
	return &Ledger{
		accounts: make(map[AccountID]*Account),
		balances: make(map[AccountID]Money),
		entries:  make(map[EntryID]*Entry),
	}
}

func (l *Ledger) CreateAccount(name string, currency Currency) (*Account, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	id := AccountID(uuid.NewString()) // Simplified for example purposes
	acct := &Account{
		ID:        id,
		Name:      name,
		Currency:  currency,
		CreatedAt: time.Now(),
	}

	l.accounts[id] = acct
	l.balances[id] = NewMoney(0, currency)
	return acct, nil
}

func (l *Ledger) GetAccount(id AccountID) (*Account, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	acct, ok := l.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return acct, nil
}

func (l *Ledger) GetBalance(accountID AccountID) (Money, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	balance, exists := l.balances[accountID]
	if !exists {
		return Money{}, ErrAccountNotFound
	}
	return balance, nil
}

func (l *Ledger) PostEntry(ctx context.Context, e *Entry) (*Entry, error) {
	if err := validateEntry(e); err != nil {
		return nil, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// ensure all accounts exist and currencies match
	for _, p := range e.Postings {
		acct, exists := l.accounts[p.AccountID]
		if !exists {
			return nil, ErrAccountNotFound
		}
		if acct.Currency != p.Amount.Currency {
			return nil, ErrCurrencyMismatch
		}
	}

	// prospective balance
	newBalances := make(map[AccountID]Money, len(l.balances))
	for id, bal := range l.balances {
		newBalances[id] = bal
	}
	for _, p := range e.Postings {
		current := newBalances[p.AccountID]
		newBalance, err := current.Add(p.Amount)
		if err != nil {
			return nil, err
		}
		newBalances[p.AccountID] = newBalance
	}

	// check non negative
	for _, bal := range newBalances {
		if bal.Amount < 0 {
			return nil, ErrNegativeBalance
		}
	}

	// assign ids and timestamps
	if e.ID == "" {
		e.ID = EntryID(uuid.NewString())
	}
	now := time.Now()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = now
	}
	if e.EffectiveAt.IsZero() {
		e.EffectiveAt = now
	}

	for id, bal := range newBalances {
		l.balances[id] = bal
	}
	l.entries[e.ID] = e

	return e, nil
}

func (l *Ledger) ListEntries(accountID AccountID) []Entry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var entries []Entry
	for _, e := range l.entries {
		if accountID == "" {
			entries = append(entries, *e)
			continue
		}
		for _, p := range e.Postings {
			if p.AccountID == accountID {
				entries = append(entries, *e)
				break
			}
		}
	}
	return entries
}

func validateEntry(e *Entry) error {
	if len(e.Postings) == 0 {
		return ErrEmptyPostings
	}

	sumByCurrency := make(map[Currency]int64)
	for _, p := range e.Postings {
		if p.Amount.isZero() {
			return ErrZeroAmountPosting
		}
		sumByCurrency[p.Amount.Currency] += p.Amount.Amount
	}

	for _, sum := range sumByCurrency {
		if sum != 0 {
			return ErrUnbalancedEntry
		}
	}

	return nil
}
