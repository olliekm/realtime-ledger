package service

import (
	"context"

	"github.com/olliekm/realtime-ledger/internal/ledger"
)

type LedgerService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*ledger.Account, error)
	GetAccount(ctx context.Context, id string) (*ledger.Account, error)
	GetBalance(ctx context.Context, id string) (ledger.Money, error)

	PostJournal(ctx context.Context, req PostJournalRequest) (*ledger.Entry, error)
	ListEntries(ctx context.Context, filter ListEntriesFilter) ([]*ledger.Entry, error)
}
