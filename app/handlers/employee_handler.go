package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/RuhullahReza/Employee-App/app/domain"
	"github.com/RuhullahReza/Employee-App/app/repositories"
	"github.com/RuhullahReza/Employee-App/app/usecases"
	"github.com/RuhullahReza/Employee-App/pkg/logger"
	"github.com/RuhullahReza/Employee-App/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	employeeUsecase usecases.EmployeeUsecase
}

var (
	validOrder = map[string]bool{
		"id":         true,
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"hire_date":  true,
		"created_at":  true,
		"updated_at":  true,
	}

	validSort = map[string]bool{
		"ASC":  true,
		"DESC": true,
	}
)

func NewEmployeeHandler(uc usecases.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUsecase: uc,
	}
}

func (h *EmployeeHandler) CreateNewEmployee(ctx *fiber.Ctx) error {
	var request domain.EmployeeRequest
	if err := ctx.BodyParser(&request); err != nil {
		logger.Log.Error(err, "failed to parse request")
		return utils.ResponseBadRequest(ctx, err.Error())
	}

	if err := utils.ValidateAndSanitizeRequest(&request); err != nil {
		logger.Log.Error(err, "body request validation error")
		return utils.ResponseBadRequest(ctx, err.Error())
	}

	res, err := h.employeeUsecase.CreateEmployee(request)
	if err != nil {
		logger.Log.Error(err, "failed to create employee")
		if errors.Is(err, usecases.ErrInvalidDate) || errors.Is(err, usecases.ErrDuplicateEmail) {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		return utils.ResponseInternalServerError(ctx, err.Error())
	}

	return utils.ResponseCreated(ctx, "Successfully create new employee", res)
}

func (h *EmployeeHandler) FindAllEmployee(ctx *fiber.Ctx) error {
	var err error

	pageNumStr := ctx.Query("pageNum")
	pageNum := 1

	if pageNumStr != "" {
		pageNum, err = strconv.Atoi(pageNumStr)
		if err != nil {
			logger.Log.Error(err, "failed to parse pageNum")
			return utils.ResponseBadRequest(ctx, "invalid page num")
		}
	}

	pageSizeStr := ctx.Query("pageSize")
	pageSize := 20

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			logger.Log.Error(err, "failed to parse pageSize")
			return utils.ResponseBadRequest(ctx, "invalid page size")
		}
	}

	orderBy := ctx.Query("orderBy")
	if _, ok := validOrder[orderBy]; !ok {
		orderBy = "created_at"
	}

	sort := ctx.Query("sort")
	if _, ok := validSort[sort]; !ok {
		sort = "DESC"
	}

	employees, err := h.employeeUsecase.GetAllEmployee(pageNum, pageSize, orderBy, sort)
	if err != nil {
		logger.Log.Error(err, "failed to get all employee")
		return utils.ResponseInternalServerError(ctx, err.Error())
	}

	return utils.ResponseOK(ctx, "Successfully get all employee data", employees)
}

func (h *EmployeeHandler) FindEmployeeById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Log.Error(err, "failed to parse id")
		return utils.ResponseBadRequest(ctx, "invalid id")
	}

	if intId < 1 {
		logger.Log.Error(err, "invalid id")
		return utils.ResponseBadRequest(ctx, "invalid id")
	}

	uintId := uint(intId)
	employee, err := h.employeeUsecase.GetEmployeeById(uintId)
	if err != nil {
		logger.Log.Error(err, "failed to get employee by id")

		if errors.Is(err, repositories.ErrRecordNotFound) {
			errMsg := fmt.Sprintf("employee with id %d not found", uintId)
			return utils.ResponseNotFound(ctx, errMsg)
		}

		return utils.ResponseInternalServerError(ctx, err.Error())
	}

	msg := fmt.Sprintf("Successfully get data for employee id %d", uintId)
	return utils.ResponseOK(ctx, msg, employee)
}

func (h *EmployeeHandler) UpdateEmployeeById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return utils.ResponseBadRequest(ctx, "invalid id")
	}

	if intId < 1 {
		return utils.ResponseBadRequest(ctx, "invalid id")
	}
	uintId := uint(intId)

	var request domain.EmployeeRequest
	if err := ctx.BodyParser(&request); err != nil {
		logger.Log.Error(err, "failed to parse body request")
		return utils.ResponseBadRequest(ctx, err.Error())
	}

	if err := utils.ValidateAndSanitizeRequest(&request); err != nil {
		logger.Log.Error(err, "body request validation error")
		return utils.ResponseBadRequest(ctx, err.Error())
	}

	res, err := h.employeeUsecase.UpdateEmployeeById(uintId, request)
	if err != nil {
		logger.Log.Error(err, "failed to update employee by id")

		if errors.Is(err, usecases.ErrInvalidDate) || errors.Is(err, usecases.ErrDuplicateEmail) {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		if errors.Is(err, repositories.ErrRecordNotFound) {
			errMsg := fmt.Sprintf("employee with id %d not found", uintId)
			return utils.ResponseNotFound(ctx, errMsg)
		}

		return utils.ResponseInternalServerError(ctx, err.Error())
	}

	msg := fmt.Sprintf("Successfully update data for employee id %d", uintId)
	return utils.ResponseOK(ctx, msg, res)
}

func (h *EmployeeHandler) DeleteEmployeeById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return utils.ResponseBadRequest(ctx, "invalid id")
	}

	uintId := uint(intId)
	if uintId < 1 {
		return utils.ResponseBadRequest(ctx, "invalid id")
	}

	err = h.employeeUsecase.DeleteEmployeeById(uintId)
	if err != nil {
		logger.Log.Error(err, "failed to delete employee by id")

		if errors.Is(err, repositories.ErrRecordNotFound) {
			errMsg := fmt.Sprintf("employee with id %d not found", uintId)
			return utils.ResponseNotFound(ctx, errMsg)
		}

		return utils.ResponseInternalServerError(ctx, err.Error())
	}

	msg := fmt.Sprintf("Successfully delete data for employee id %d", uintId)
	return utils.ResponseOK(ctx, msg, nil)
}
