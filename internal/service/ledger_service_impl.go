package service

import (
	"context"

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
	if err != nil {
		return nil, err
	}

	if err := s.store.CreateAccount(ctx, acct); err != nil {
		return nil, err
	}
	return acct, nil
}
