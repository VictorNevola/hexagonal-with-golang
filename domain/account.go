package domain

import (
	"math/rand"
	"time"
)

type (
	Account struct {
		ID            uint64    `db:"id"`
		AccountNumber uint64    `db:"account_number"`
		CustomerID    string    `db:"customer_id"`
		OpeningDate   time.Time `db:"opening_date"`
		AccountType   string    `db:"account_type"`
		Amount        float64   `db:"amount"`
		Status        string    `db:"status"`
	}

	AccountRepositoryPort interface {
		Save(account Account) (*Account, error)
	}
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func (acc *Account) GenerateAccountNumber() {
	acc.AccountNumber = uint64(100000 + rng.Intn(900000))
}

func (acc *Account) ValidateMinValueToOpenAccount() bool {
	return acc.Amount >= 5000
}
