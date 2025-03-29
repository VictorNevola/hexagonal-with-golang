package database

import "github.com/VictorNevola/hexagonal/domain"

type (
	CustomerRepositoryMock struct {
		Customers []domain.Customer
	}
)

func (s CustomerRepositoryMock) FindAll() ([]domain.Customer, error) {
	return s.Customers, nil
}

func (s CustomerRepositoryMock) ByID(id string) (*domain.Customer, error) {
	for _, customer := range s.Customers {
		if customer.ID == id {
			return &customer, nil
		}
	}
	return nil, nil
}

func NewCustomerRespositoryMock() CustomerRepositoryMock {
	return CustomerRepositoryMock{
		Customers: []domain.Customer{
			{
				ID:          "1001",
				Name:        "John Doe",
				Zipcode:     "12345",
				DateOfBirth: "01-01-1970",
				Status:      "active",
			},
			{
				ID:          "1002",
				Name:        "Jane Doe",
				Zipcode:     "12345",
				DateOfBirth: "01-01-1970",
				Status:      "active",
			},
		},
	}
}
