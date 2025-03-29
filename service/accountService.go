package service

import (
	"time"

	"github.com/VictorNevola/hexagonal/domain"
)

type (
	AccountCreateDTO struct {
		CustomerID  string  `json:"customer_id"`
		AccountType string  `json:"account_type"`
		Amount      float64 `json:"amount"`
	}

	AccountResponseDTO struct {
		AccountID     uint64 `json:"account_id,omitempty"`
		AccountNumber uint64 `json:"account_number,omitempty"`
		CustomerID    string `json:"customer_id,omitempty"`
	}

	AccountServicePort interface {
		NewAccount(account AccountCreateDTO) (AccountResponseDTO, error)
	}

	AcountServiceAdapter struct {
		repo domain.AccountRepositoryPort
	}
)

func NewAccountService(repo domain.AccountRepositoryPort) AcountServiceAdapter {
	return AcountServiceAdapter{
		repo: repo,
	}
}

func (s AcountServiceAdapter) NewAccount(account AccountCreateDTO) (AccountResponseDTO, error) {
	newAccount := domain.Account{
		CustomerID:  account.CustomerID,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		OpeningDate: time.Now(),
		Status:      "active",
	}

	if !newAccount.ValidateMinValueToOpenAccount() {
		return AccountResponseDTO{}, ErrorAccountCannotBeOpened
	}

	existingAccount, err := s.repo.GetAccountByCustomerID(newAccount.CustomerID)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	if existingAccount != nil {
		return AccountResponseDTO{}, ErrorAccountAlreadyExists
	}

	newAccount.GenerateAccountNumber()

	createdAccount, err := s.repo.Save(newAccount)
	if err != nil {
		return AccountResponseDTO{}, err
	}

	return AccountResponseDTO{
		AccountID:     createdAccount.ID,
		AccountNumber: createdAccount.AccountNumber,
		CustomerID:    createdAccount.CustomerID,
	}, nil
}
