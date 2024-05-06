package app

import (
	"github.com/RuhullahReza/Employee-App/app/handlers"
	"github.com/RuhullahReza/Employee-App/app/repositories"
	"github.com/RuhullahReza/Employee-App/app/usecases"
	config "github.com/RuhullahReza/Employee-App/config"
	"github.com/RuhullahReza/Employee-App/pkg/database"
	"github.com/RuhullahReza/Employee-App/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func NewServer(cfg *config.Config) *fiber.App {
	db := database.Init(cfg)
	database.AutoMigrate(db)

	employeeRepository := repositories.NewEmployeeRepository(db)
	empolyeeUsecase := usecases.NewEmployeeUsecase(employeeRepository)
	employeeHandler := handlers.NewEmployeeHandler(empolyeeUsecase)

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	router := routes.NewRoutes(app, employeeHandler)
	router.Init(cfg.EndpointPrefix)

	return app
}
