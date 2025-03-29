package service

import (
	"database/sql"
	"errors"

	"github.com/VictorNevola/hexagonal/domain"
)

type (
	CustomerServicePort interface {
		GetAllCustomers() ([]domain.Customer, error)
		GetCustomer(id string) (*domain.Customer, error)
	}

	CustomerServiceAdapter struct {
		repo domain.CustomerRepositoryPort
	}
)

func NewCustomerService(repo domain.CustomerRepositoryPort) CustomerServiceAdapter {
	return CustomerServiceAdapter{
		repo: repo,
	}
}

func (s CustomerServiceAdapter) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s CustomerServiceAdapter) GetCustomer(id string) (*domain.Customer, error) {
	customer, err := s.repo.ByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCustomerNotFound
	}

	if err != nil {
		return nil, err
	}

	return customer, nil
}
