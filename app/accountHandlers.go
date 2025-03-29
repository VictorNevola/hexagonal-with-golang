package app

import (
	"net/http"

	"github.com/VictorNevola/hexagonal/service"
	"github.com/gofiber/fiber/v2"
)

type (
	AccountHandler struct {
		service service.AccountServicePort
	}
)

func NewAccountHandler(app *fiber.App, service service.AccountServicePort) {
	httpHandler := AccountHandler{
		service: service,
	}

	app.Route("/accounts", func(r fiber.Router) {
		r.Post("/", httpHandler.CreateAccount)
	})
}

func (h AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var request service.AccountCreateDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	account, err := h.service.NewAccount(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(account)
}
