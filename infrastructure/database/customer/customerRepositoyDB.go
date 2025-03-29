package database

import (
	"github.com/VictorNevola/hexagonal/domain"
	"github.com/VictorNevola/hexagonal/logger"
	"github.com/jmoiron/sqlx"
)

type (
	CustomerRepositoryAdapter struct {
		Db *sqlx.DB
	}
)

func NewCustomerRepository(db *sqlx.DB) CustomerRepositoryAdapter {
	repo := CustomerRepositoryAdapter{Db: db}

	repo.createTableIfNotExists()

	return repo
}

func (repo *CustomerRepositoryAdapter) createTableIfNotExists() {
	query := `CREATE TABLE IF NOT EXISTS customers (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255),
		zipcode VARCHAR(10),
		date_of_birth DATE,
		status VARCHAR(10)
	)`

	_, err := repo.Db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (r CustomerRepositoryAdapter) FindAll() ([]domain.Customer, error) {
	var customers []domain.Customer

	err := r.Db.Select(&customers, "SELECT id, name, zipcode, date_of_birth, status FROM customers")
	if err != nil {
		logger.Error(CustomerNotFound, CustomerNotFoundError(err))
		return nil, err
	}

	return customers, nil
}

func (r CustomerRepositoryAdapter) ByID(id string) (*domain.Customer, error) {
	var c domain.Customer

	err := r.Db.Get(&c, "SELECT id, name, zipcode, date_of_birth, status FROM customers WHERE id = ?", id)
	if err != nil {
		logger.Error(CustomerNotFound, CustomerNotFoundError(err))
		return nil, err
	}

	return &c, nil
}
