package app

import (
	"errors"

	"github.com/VictorNevola/hexagonal/service"
	"github.com/gofiber/fiber/v2"
)

type (
	CustomerHandler struct {
		service service.CustomerServicePort
	}

	CustomerResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Zipcode     string `json:"zipcode"`
		DateOfBirth string `json:"date_of_birth"`
		Status      string `json:"status"`
	}
)

func NewCustomerHandler(app *fiber.App, service service.CustomerServicePort) {
	httpHandler := CustomerHandler{
		service: service,
	}

	app.Route("/customers", func(r fiber.Router) {
		r.Get("/", httpHandler.GetAllCustomers)
		r.Get("/:id", httpHandler.GetByID)
	})
}

func (h CustomerHandler) GetAllCustomers(c *fiber.Ctx) error {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		return err
	}

	var response []CustomerResponse
	for _, customer := range customers {
		response = append(response, CustomerResponse{
			ID:          customer.ID,
			Name:        customer.Name,
			Zipcode:     customer.Zipcode,
			DateOfBirth: customer.DateOfBirth,
			Status:      customer.Status,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h CustomerHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	customer, err := h.service.GetCustomer(id)
	if errors.Is(err, service.ErrCustomerNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err != nil {
		return err
	}

	response := CustomerResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.Status,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
