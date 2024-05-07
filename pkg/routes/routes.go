package routes

import (
	"github.com/RuhullahReza/Employee-App/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	router          fiber.Router
	employeeHandler *handlers.EmployeeHandler
}

func NewRoutes(app *fiber.App, h *handlers.EmployeeHandler) *Routes {
	return &Routes{
		router:          app,
		employeeHandler: h,
	}
}

func (r *Routes) employeeRoutes(prefix string) {
	resources := r.router.Group(prefix + "/employees")
	resources.Post("/", r.employeeHandler.CreateNewEmployee)
	resources.Get("/", r.employeeHandler.FindAllEmployee)
	resources.Get("/:id", r.employeeHandler.FindEmployeeById)
	resources.Put("/:id", r.employeeHandler.UpdateEmployeeById)
	resources.Delete("/:id", r.employeeHandler.DeleteEmployeeById)
}

func (r *Routes) Init(prefix string) {
	r.employeeRoutes(prefix)
}
