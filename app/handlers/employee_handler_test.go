package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RuhullahReza/Employee-App/app/domain"
	"github.com/RuhullahReza/Employee-App/app/mocks"
	"github.com/RuhullahReza/Employee-App/app/repositories"
	"github.com/RuhullahReza/Employee-App/app/usecases"
	"github.com/RuhullahReza/Employee-App/pkg/logger"
	"github.com/RuhullahReza/Employee-App/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeHandler(t *testing.T) {
	logger.Init()

	uc := new(mocks.EmployeeUsecase)
	h := NewEmployeeHandler(uc)

	app := fiber.New()
	app.Post("api/employee", h.CreateNewEmployee)
	app.Get("api/employee", h.FindAllEmployee)
	app.Get("api/employee/:id", h.FindEmployeeById)
	app.Put("api/employee/:id", h.UpdateEmployeeById)
	app.Delete("api/employee/:id", h.DeleteEmployeeById)

	t.Run("Test Create Employee SUCCESS", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		parsedTime, _ := utils.ParseDateString(req.HireDate)

		employeeData := domain.EmployeeResponse{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     req.Email,
			HireDate:  parsedTime,
		}

		uc.On("CreateEmployee", req).
			Return(employeeData, nil).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPost, "/api/employee", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Test Create Employee Internal Error", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		employeeData := domain.EmployeeResponse{}
		uc.On("CreateEmployee", req).
			Return(employeeData, errors.New("error")).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPost, "/api/employee", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Test Create Employee BAD REQUEST duplicate email", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		employeeData := domain.EmployeeResponse{}
		uc.On("CreateEmployee", req).
			Return(employeeData, usecases.ErrDuplicateEmail).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPost, "/api/employee", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Create Employee BAD REQUEST validation error", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "",
			LastName:  "",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPost, "/api/employee", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Create Employee BAD REQUEST invalid body", func(t *testing.T) {
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode("abc")

		httpReq := httptest.NewRequest(http.MethodPost, "/api/employee", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Get All Employee SUCCESS", func(t *testing.T) {
		response := domain.PaginationResponse{
			PageNum:   1,
			PageSize:  20,
			TotalPage: 1,
			Data:      []domain.EmployeeResponse{},
		}

		uc.On("GetAllEmployee", 1, 20, "created_at", "DESC").
			Return(response, nil).
			Once()

		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Test Get All Employee INTERNAL ERROR", func(t *testing.T) {
		response := domain.PaginationResponse{}

		uc.On("GetAllEmployee", 1, 20, "created_at", "DESC").
			Return(response, errors.New("error")).
			Once()

		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Test Get All Employee BAD REQUEST invalid page size", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee?pageSize=abc", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Get All Employee BAD REQUEST invalid page num", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee?pageNum=abc", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Get Employee By ID SUCCESS", func(t *testing.T) {
		parsedDate, _ := utils.ParseDateString("2024-03-03")
		response := domain.EmployeeResponse{
			FirstName: "Ozza",
			LastName:  "Reza",
			Email:     "test.test@gmail.com",
			HireDate:  parsedDate,
		}

		uc.On("GetEmployeeById", uint(1)).
			Return(response, nil).
			Once()

		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee/1", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Test Get Employee By ID INTERNAL ERROR", func(t *testing.T) {
		response := domain.EmployeeResponse{}

		uc.On("GetEmployeeById", uint(1)).
			Return(response, errors.New("error")).
			Once()

		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee/1", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Test Get Employee By ID NOT FOUND", func(t *testing.T) {
		response := domain.EmployeeResponse{}

		uc.On("GetEmployeeById", uint(1)).
			Return(response, repositories.ErrRecordNotFound).
			Once()

		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee/1", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Test Get Employee By ID BAD REQUEST invalid id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee/0", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Get Employee By ID BAD REQUEST failed to parse id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodGet, "/api/employee/abc", nil)
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID SUCCESS", func(t *testing.T) {
		id := uint(1)
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		parsedTime, _ := utils.ParseDateString(req.HireDate)

		employeeData := domain.EmployeeResponse{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     req.Email,
			HireDate:  parsedTime,
		}

		uc.On("UpdateEmployeeById", id, req).
			Return(employeeData, nil).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID INTERNAL ERROR", func(t *testing.T) {
		id := uint(1)
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		employeeData := domain.EmployeeResponse{}

		uc.On("UpdateEmployeeById", id, req).
			Return(employeeData, errors.New("error")).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID NOT FOUND", func(t *testing.T) {
		id := uint(1)
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		employeeData := domain.EmployeeResponse{}

		uc.On("UpdateEmployeeById", id, req).
			Return(employeeData, repositories.ErrRecordNotFound).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID BAD REQUEST duplicate email", func(t *testing.T) {
		id := uint(1)
		req := domain.EmployeeRequest{
			FirstName: "Reza",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		employeeData := domain.EmployeeResponse{}

		uc.On("UpdateEmployeeById", id, req).
			Return(employeeData, usecases.ErrDuplicateEmail).
			Once()

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID BAD REQUEST empty name", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "",
			LastName:  "Ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(req)

		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", &buf)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID BAD REQUEST invalid body", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/1", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID BAD REQUEST invalid id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/0", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Update Employee By ID BAD REQUEST faild to parse id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodPut, "/api/employee/abc", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Delete Employee By ID SUCCESS", func(t *testing.T) {
		uc.On("DeleteEmployeeById", uint(1)).
			Return(nil).
			Once()

		httpReq := httptest.NewRequest(http.MethodDelete, "/api/employee/1", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Test Delete Employee By ID INTERNAL ERROR", func(t *testing.T) {
		uc.On("DeleteEmployeeById", uint(1)).
			Return(errors.New("error")).
			Once()

		httpReq := httptest.NewRequest(http.MethodDelete, "/api/employee/1", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("Test Delete Employee By ID NOT FOUND", func(t *testing.T) {
		uc.On("DeleteEmployeeById", uint(1)).
			Return(repositories.ErrRecordNotFound).
			Once()

		httpReq := httptest.NewRequest(http.MethodDelete, "/api/employee/1", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("Test Delete Employee By ID BAD REQUEST invalid id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodDelete, "/api/employee/0", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Test Delete Employee By ID BAD REQUEST failed to parse id", func(t *testing.T) {
		httpReq := httptest.NewRequest(http.MethodDelete, "/api/employee/abc", nil)
		httpReq.Header.Set("content-type", "application/json")
		resp, err := app.Test(httpReq, 2)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
