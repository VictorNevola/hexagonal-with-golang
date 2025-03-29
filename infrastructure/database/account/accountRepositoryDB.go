package account

import (
	"github.com/VictorNevola/hexagonal/domain"
	"github.com/jmoiron/sqlx"
)

type (
	AccountRepositoryAdapter struct {
		DB *sqlx.DB
	}
)

func NewAccountRepository(db *sqlx.DB) AccountRepositoryAdapter {
	repo := AccountRepositoryAdapter{DB: db}

	repo.createTableIfNotExists()

	return repo
}

func (repo *AccountRepositoryAdapter) createTableIfNotExists() {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		account_number BIGINT UNSIGNED UNIQUE NOT NULL,
		customer_id VARCHAR(36) NOT NULL,
		opening_date DATETIME NOT NULL,
		account_type VARCHAR(10) NOT NULL,
		amount DECIMAL(10,2) NOT NULL DEFAULT 0.00,
		status VARCHAR(10) NOT NULL
	)`

	_, err := repo.DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (r AccountRepositoryAdapter) Save(account domain.Account) (*domain.Account, error) {
	query := `INSERT INTO accounts (account_number, customer_id, opening_date, account_type, amount, status)
		VALUES (?, ?, ?, ?, ?)`

	result, err := r.DB.Exec(query, account.AccountNumber, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	account.ID = uint64(id)

	return &account, nil
}
