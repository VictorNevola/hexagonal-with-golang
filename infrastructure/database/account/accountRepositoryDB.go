package account

import (
	"time"

	"github.com/VictorNevola/hexagonal/domain"
	"github.com/jmoiron/sqlx"
)

type (
	AccountRepositoryAdapter struct {
		DB *sqlx.DB
	}

	AccountDTO struct {
		ID            uint64  `db:"id"`
		AccountNumber uint64  `db:"account_number"`
		CustomerID    string  `db:"customer_id"`
		OpeningDate   string  `db:"opening_date"`
		AccountType   string  `db:"account_type"`
		Amount        float64 `db:"amount"`
		Status        string  `db:"status"`
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
		VALUES (?, ?, ?, ?, ?, ?)`

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

func (r AccountRepositoryAdapter) GetAccountByCustomerID(customerID string) (*domain.Account, error) {
	query := `SELECT * FROM accounts WHERE customer_id = ?`

	var accountDTO AccountDTO
	err := r.DB.Get(&accountDTO, query, customerID)
	if err != nil {
		return nil, err
	}

	openingDate, _ := time.Parse("2006-01-02 15:04:05", accountDTO.OpeningDate)

	account := domain.Account{
		ID:            accountDTO.ID,
		AccountNumber: accountDTO.AccountNumber,
		CustomerID:    accountDTO.CustomerID,
		OpeningDate:   openingDate,
		AccountType:   accountDTO.AccountType,
		Amount:        accountDTO.Amount,
		Status:        accountDTO.Status,
	}

	return &account, nil
}
