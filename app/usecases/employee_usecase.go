package usecases

import (
	"errors"

	"github.com/RuhullahReza/Employee-App/app/domain"
	"github.com/RuhullahReza/Employee-App/app/repositories"
	"github.com/RuhullahReza/Employee-App/pkg/logger"
	"github.com/RuhullahReza/Employee-App/pkg/utils"
)

type EmployeeUsecase interface {
	CreateEmployee(req domain.EmployeeRequest) (domain.EmployeeResponse, error)
	GetAllEmployee(page, limit int, orderBy, sort string) (domain.PaginationResponse, error)
	GetEmployeeById(id uint) (domain.EmployeeResponse, error)
	UpdateEmployeeById(id uint, req domain.EmployeeRequest) (domain.EmployeeResponse, error)
	DeleteEmployeeById(id uint) error
}

type employeeUsecase struct {
	employeeRepository repositories.EmployeeRepository
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrInvalidDate    = errors.New("invalid date format")
)

func NewEmployeeUsecase(employeeRepository repositories.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{
		employeeRepository: employeeRepository,
	}
}

func (uc *employeeUsecase) CreateEmployee(req domain.EmployeeRequest) (domain.EmployeeResponse, error) {
	parsedDate, err := utils.ParseDateString(req.HireDate)
	if err != nil {
		logger.Log.Error(err, "failed to parse Hire Date")
		return domain.EmployeeResponse{}, ErrInvalidDate
	}

	foundEmployee, err := uc.employeeRepository.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, repositories.ErrRecordNotFound) {
		logger.Log.Error(err, "failed to find employee by email")
		return domain.EmployeeResponse{}, err
	}

	if foundEmployee.ID != 0 {
		return domain.EmployeeResponse{}, ErrDuplicateEmail
	}

	newEmployee := domain.Employee{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		HireDate:  parsedDate,
	}

	if err := uc.employeeRepository.Store(&newEmployee); err != nil {
		logger.Log.Error(err, "failed to store new employee data")
		return domain.EmployeeResponse{}, err
	}

	res := domain.EmployeeResponse{
		Id:        newEmployee.ID,
		FirstName: newEmployee.FirstName,
		LastName:  newEmployee.LastName,
		Email:     newEmployee.Email,
		HireDate:  newEmployee.HireDate,
		CreatedAt: newEmployee.CreatedAt,
		UpdatedAt: newEmployee.UpdatedAt,
	}

	logger.Log.Info("successfully create employee with id : ", newEmployee.ID)
	return res, nil
}

func (uc *employeeUsecase) GetAllEmployee(page, limit int, orderBy, sort string) (domain.PaginationResponse, error) {
	offset := (page - 1) * limit
	employees, count, err := uc.employeeRepository.FindAll(limit, offset, orderBy, sort)
	if err != nil {
		logger.Log.Error(err, "failed to find all employee data")
		return domain.PaginationResponse{}, err
	}

	var employeeResponses []domain.EmployeeResponse
	for _, e := range employees {
		employee := domain.EmployeeResponse{
			Id:        e.ID,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			Email:     e.Email,
			HireDate:  e.HireDate,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		}

		employeeResponses = append(employeeResponses, employee)
	}

	totalPage := count / int64(limit)
	if count%int64(limit) != 0 {
		totalPage++
	}

	res := domain.PaginationResponse{
		PageNum:   page,
		PageSize:  limit,
		TotalPage: totalPage,
		Data:      employeeResponses,
	}

	return res, nil
}

func (uc *employeeUsecase) GetEmployeeById(id uint) (domain.EmployeeResponse, error) {
	employee, err := uc.employeeRepository.FindById(id)
	if err != nil {
		logger.Log.Error(err, "failed to find employee by id")
		return domain.EmployeeResponse{}, err
	}

	employeeResponse := domain.EmployeeResponse{
		Id:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		HireDate:  employee.HireDate,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}

	return employeeResponse, nil
}

func (uc *employeeUsecase) UpdateEmployeeById(id uint, req domain.EmployeeRequest) (domain.EmployeeResponse, error) {
	parsedDate, err := utils.ParseDateString(req.HireDate)
	if err != nil {
		logger.Log.Error(err, "failed to parse Hire Date")
		return domain.EmployeeResponse{}, ErrInvalidDate
	}

	_, err = uc.employeeRepository.FindById(id)
	if err != nil {
		logger.Log.Error(err, "failed to find employee by id")
		return domain.EmployeeResponse{}, err
	}

	foundEmployee, err := uc.employeeRepository.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, repositories.ErrRecordNotFound) {
		logger.Log.Error(err, "failed to find employee by email")
		return domain.EmployeeResponse{}, err
	}

	if foundEmployee.ID != 0 && foundEmployee.ID != id {
		return domain.EmployeeResponse{}, ErrDuplicateEmail
	}

	updatedEmployee := domain.Employee{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		HireDate:  parsedDate,
	}

	if err := uc.employeeRepository.UpdateById(&updatedEmployee); err != nil {
		logger.Log.Error(err, "failed to update employee by id")
		return domain.EmployeeResponse{}, err
	}

	res := domain.EmployeeResponse{
		Id:        updatedEmployee.ID,
		FirstName: updatedEmployee.FirstName,
		LastName:  updatedEmployee.LastName,
		Email:     updatedEmployee.Email,
		HireDate:  updatedEmployee.HireDate,
		UpdatedAt: updatedEmployee.UpdatedAt,
	}

	logger.Log.Info("successfully update employee with id : ", updatedEmployee.ID)
	return res, nil
}

func (uc *employeeUsecase) DeleteEmployeeById(id uint) error {
	err := uc.employeeRepository.DeleteById(id)
	if err != nil {
		logger.Log.Error(err, "failed to delete employee by id")
		return err
	}

	logger.Log.Info("successfully delete employee with id : ", id)
	return nil
}
