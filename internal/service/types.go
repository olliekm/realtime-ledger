package service

import "time"

type CreateAccountRequest struct {
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Description string `json:"description,omitempty"`
}

type PostJournalEntry struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
	Side      string `json:"side"`
	Memo      string `json:"memo,omitempty"`
}

type PostJournalRequest struct {
	ExternalID string             `json:"external_id,omitempty"`
	At         time.Time          `json:"at"`
	Entries    []PostJournalEntry `json:"entries"`
}

type ListEntriesFilter struct {
	AccountID string
	From      *time.Time
	To        *time.Time
	Limit     int
	Offset    int
}
