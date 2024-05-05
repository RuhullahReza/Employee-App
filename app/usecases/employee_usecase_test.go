package usecases

import (
	"errors"
	"testing"

	"github.com/RuhullahReza/Employee-App/app/domain"
	"github.com/RuhullahReza/Employee-App/app/mocks"
	"github.com/RuhullahReza/Employee-App/app/repositories"
	"github.com/RuhullahReza/Employee-App/pkg/logger"
	"github.com/RuhullahReza/Employee-App/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	er := mocks.NewEmployeeRepository(t)
	uc := NewEmployeeUsecase(er)
	logger.Init()

	req := domain.EmployeeRequest{
		FirstName: "reza",
		LastName:  "ozza",
		Email:     "test.test@gmail.com",
		HireDate:  "2024-03-03",
	}

	parsedDate, _ := utils.ParseDateString(req.HireDate)

	newEmployee := domain.Employee{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		HireDate:  parsedDate,
	}

	t.Run("success", func(t *testing.T) {
		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, repositories.ErrRecordNotFound).
			Once()

		er.On("Store", &newEmployee).
			Return(nil).
			Once()

		res, err := uc.CreateEmployee(req)
		assert.NoError(t, err)

		assert.Equal(t, newEmployee.FirstName, res.FirstName)
		assert.Equal(t, newEmployee.LastName, res.LastName)
		assert.Equal(t, newEmployee.Email, res.Email)
		assert.Equal(t, newEmployee.HireDate, res.HireDate)
	})

	t.Run("failed on store", func(t *testing.T) {
		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, repositories.ErrRecordNotFound).
			Once()

		er.On("Store", &newEmployee).
			Return(errors.New("error")).
			Once()

		_, err := uc.CreateEmployee(req)
		assert.Error(t, err)
	})

	t.Run("duplicate email", func(t *testing.T) {
		er.On("FindByEmail", req.Email).
			Return(domain.Employee{ID: 1}, repositories.ErrRecordNotFound).
			Once()

		_, err := uc.CreateEmployee(req)
		assert.ErrorIs(t, err, ErrDuplicateEmail)
	})

	t.Run("failed to find email", func(t *testing.T) {
		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, errors.New("error")).
			Once()

		_, err := uc.CreateEmployee(req)
		assert.Error(t, err)
	})

	t.Run("failed to parse date", func(t *testing.T) {
		req.HireDate = "123-23-21"

		_, err := uc.CreateEmployee(req)
		assert.ErrorIs(t, err, ErrInvalidDate)
	})
}

func TestGetAllEmployee(t *testing.T) {
	er := mocks.NewEmployeeRepository(t)
	uc := NewEmployeeUsecase(er)
	logger.Init()

	parsedDate, _ := utils.ParseDateString("2024-03-03")

	newEmployee := domain.Employee{
		FirstName: "abc",
		LastName:  "def",
		Email:     "abc.def@gmail.com",
		HireDate:  parsedDate,
	}

	t.Run("success", func(t *testing.T) {
		er.On("FindAll", 20, 0, "id", "ASC").
			Return([]domain.Employee{newEmployee}, int64(1), nil).
			Once()

		res, err := uc.GetAllEmployee(1, 20, "id", "ASC")
		assert.NoError(t, err)

		assert.Equal(t, res.Data.([]domain.EmployeeResponse)[0].FirstName, newEmployee.FirstName)
		assert.Equal(t, res.Data.([]domain.EmployeeResponse)[0].LastName, newEmployee.LastName)
		assert.Equal(t, res.Data.([]domain.EmployeeResponse)[0].Email, newEmployee.Email)
		assert.Equal(t, res.Data.([]domain.EmployeeResponse)[0].HireDate, newEmployee.HireDate)
	})

	t.Run("failed to find all", func(t *testing.T) {
		er.On("FindAll", 20, 0, "id", "ASC").
			Return(nil, int64(-1), errors.New("error")).
			Once()

		_, err := uc.GetAllEmployee(1, 20, "id", "ASC")
		assert.Error(t, err)
	})
}

func TestGetEmployeeById(t *testing.T) {
	er := mocks.NewEmployeeRepository(t)
	uc := NewEmployeeUsecase(er)
	logger.Init()

	parsedDate, _ := utils.ParseDateString("2024-03-03")

	newEmployee := domain.Employee{
		FirstName: "abc",
		LastName:  "def",
		Email:     "abc.def@gmail.com",
		HireDate:  parsedDate,
	}

	id := uint(1)

	t.Run("success", func(t *testing.T) {
		er.On("FindById", id).
			Return(newEmployee, nil).
			Once()

		res, err := uc.GetEmployeeById(id)
		assert.NoError(t, err)

		assert.Equal(t, newEmployee.FirstName, res.FirstName)
		assert.Equal(t, newEmployee.LastName, res.LastName)
		assert.Equal(t, newEmployee.Email, res.Email)
		assert.Equal(t, newEmployee.HireDate, res.HireDate)
	})

	t.Run("failed to find", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{}, errors.New("error")).
			Once()

		_, err := uc.GetEmployeeById(id)
		assert.Error(t, err)
	})
}

func TestUpdateEmployeeById(t *testing.T) {
	er := mocks.NewEmployeeRepository(t)
	uc := NewEmployeeUsecase(er)
	logger.Init()

	id := uint(1)
	req := domain.EmployeeRequest{
		FirstName: "reza",
		LastName:  "ozza",
		Email:     "test.test@gmail.com",
		HireDate:  "2024-03-03",
	}

	parsedDate, _ := utils.ParseDateString(req.HireDate)

	updated := domain.Employee{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		HireDate:  parsedDate,
	}

	t.Run("success", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{ID: id}, nil).
			Once()

		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, nil).
			Once()

		er.On("UpdateById", &updated).
			Return(nil).
			Once()

		res, err := uc.UpdateEmployeeById(id, req)
		assert.NoError(t, err)

		assert.Equal(t, updated.FirstName, res.FirstName)
		assert.Equal(t, updated.LastName, res.LastName)
		assert.Equal(t, updated.Email, res.Email)
		assert.Equal(t, updated.HireDate, res.HireDate)
	})

	t.Run("fail to update by id", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{ID: id}, nil).
			Once()

		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, nil).
			Once()

		er.On("UpdateById", &updated).
			Return(errors.New("error")).
			Once()

		_, err := uc.UpdateEmployeeById(id, req)
		assert.Error(t, err)
	})

	t.Run("duplicate email", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{ID: id}, nil).
			Once()

		er.On("FindByEmail", req.Email).
			Return(domain.Employee{ID: 2}, nil).
			Once()

		_, err := uc.UpdateEmployeeById(id, req)
		assert.ErrorIs(t, err, ErrDuplicateEmail)
	})

	t.Run("fail to find by email", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{ID: id}, nil).
			Once()

		er.On("FindByEmail", req.Email).
			Return(domain.Employee{}, errors.New("error")).
			Once()

		_, err := uc.UpdateEmployeeById(id, req)
		assert.Error(t, err)
	})

	t.Run("fail to find by id", func(t *testing.T) {
		er.On("FindById", id).
			Return(domain.Employee{}, errors.New("error")).
			Once()

		_, err := uc.UpdateEmployeeById(id, req)
		assert.Error(t, err)
	})

	t.Run("fail to parse date", func(t *testing.T) {
		req.HireDate = "abc"

		_, err := uc.UpdateEmployeeById(id, req)
		assert.Error(t, err)
	})
}

func TestDeleteEmployeeById(t *testing.T) {
	er := mocks.NewEmployeeRepository(t)
	uc := NewEmployeeUsecase(er)
	logger.Init()

	id := uint(1)

	t.Run("success", func(t *testing.T) {
		er.On("DeleteById", id).
			Return(nil).
			Once()

		err := uc.DeleteEmployeeById(id)
		assert.NoError(t, err)
	})

	t.Run("fail to delete", func(t *testing.T) {
		er.On("DeleteById", id).
			Return(errors.New("error")).
			Once()

		err := uc.DeleteEmployeeById(id)
		assert.Error(t, err)
	})
}
