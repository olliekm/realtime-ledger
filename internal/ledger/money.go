package ledger

import "fmt"

type Curency string

const (
	USD Curency = "USD"
	EUR Curency = "EUR"
	GBP Curency = "GBP"
)

type Money struct {
	Amount   int64
	Currency Curency
}

func NewMoney(amount int64, c Curency) Money {
	return Money{Amount: amount, Currency: c}
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}

func (m Money) sameCurrency(other Money) bool {
	return m.Currency == other.Currency
}

func (m Money) Add(other Money) (Money, error) {
	if !m.sameCurrency(other) {
		return Money{}, fmt.Errorf("cannot add money with different currencies: %s and %s", m.Currency, other.Currency)
	}
	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Negate() Money {
	return Money{
		Amount:   -m.Amount,
		Currency: m.Currency,
	}
}

func (m Money) isZero() bool {
	return m.Amount == 0
}
