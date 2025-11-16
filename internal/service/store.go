package service

import (
	"context"
	"time"

	"github.com/olliekm/realtime-ledger/internal/ledger"
)

type LedgerStore interface {
	CreateAccount(ctx context.Context, a *ledger.Account) error
	GetAccount(ctx context.Context, id string) (*ledger.Account, error)

	InsertJournal(ctx context.Context, j *ledger.Journal) error
	ListEntries(ctx context.Context, accountID string, from, to *time.Time, limit, offset int) ([]ledger.Entry, error)

	GetBalance(ctx context.Context, accountID string) (ledger.Money, error)
}
