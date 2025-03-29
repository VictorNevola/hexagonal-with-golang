package app

import (
	"github.com/VictorNevola/hexagonal/config"
	accountRepo "github.com/VictorNevola/hexagonal/infrastructure/database/account"
	customerRepo "github.com/VictorNevola/hexagonal/infrastructure/database/customer"
	"github.com/VictorNevola/hexagonal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	sqlConnection := config.ConnectToMySql("admin:admin@/HEXAGONAL_ARCHITECTURE")
	app := fiber.New()
	app.Use(logger.New())

	// Repositories
	customersRepo := customerRepo.NewCustomerRepository(sqlConnection)
	accountRepo := accountRepo.NewAccountRepository(sqlConnection)

	// Services
	customersService := service.NewCustomerService(customersRepo)
	accountService := service.NewAccountService(accountRepo)

	// Routers
	NewCustomerHandler(app, customersService)
	NewAccountHandler(app, accountService)

	app.Listen(":3000")
}
