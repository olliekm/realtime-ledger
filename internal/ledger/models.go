package ledger

import "time"

type AccountID string
type EntryID string

// Represents a ledger account.
type Account struct {
	ID        AccountID
	Name      string
	Currency  Curency
	CreatedAt time.Time
}

// Posting is a single line in a journal entry.
type Posting struct {
	AccountID AccountID
	Amount    Money
}

// Entry is a journal entry with multiple postings.
type Entry struct {
	ID          EntryID
	EffectiveAt time.Time
	CreatedAt   time.Time
	Postings    []Posting
	Metadata    map[string]string
}

// BalanceSnapshot represents the balance of an account at a specific point in time.
type BalanceSnapshot struct {
	AccountID AccountID
	Balance   Money
	AsOf      time.Time
	Version   int64
}
