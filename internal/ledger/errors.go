package ledger

import "errors"

var (
	ErrCurrencyMismatch  = errors.New("currency mismatch")
	ErrUnbalancedEntry   = errors.New("unbalanced entry")
	ErrEmptyPostings     = errors.New("entry has no postings")
	ErrAccountNotFound   = errors.New("account not found")
	ErrZeroAmountPosting = errors.New("posting has zero amount")
	ErrNegativeBalance   = errors.New("operation would result in negative balance")
)
