package domain

type (
	Customer struct {
		ID          string `db:"id"`
		Name        string `db:"name"`
		Zipcode     string `db:"zipcode"`
		DateOfBirth string `db:"date_of_birth"`
		Status      string `db:"status"`
	}

	CustomerRepositoryPort interface {
		FindAll() ([]Customer, error)
		ByID(id string) (*Customer, error)
	}
)
