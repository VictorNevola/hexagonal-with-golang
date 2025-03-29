package service

import "errors"

var (
	ErrCustomerNotFound        = errors.New("customer not found")
	ErrorAccountCannotBeOpened = errors.New("account cannot be opened, the minimum value is 5000")
	ErrorAccountAlreadyExists  = errors.New("account already exists for this customer")
)
