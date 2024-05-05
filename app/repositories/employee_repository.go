package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/RuhullahReza/Employee-App/app/domain"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Store(employee *domain.Employee) error
	FindAll(limit, offset int, orderBy, sort string) ([]domain.Employee, int64, error)
	FindById(id uint) (domain.Employee, error)
	FindByEmail(email string) (domain.Employee, error)
	UpdateById(employee *domain.Employee) error
	DeleteById(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

var (
	ErrNilReference   = errors.New("invalid data, nil struct")
	ErrRecordNotFound = errors.New("record not found")
)

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) Store(employee *domain.Employee) error {
	if employee == nil {
		return ErrNilReference
	}

	created := r.db.Create(employee)
	if created.Error != nil {
		return created.Error
	}

	return nil
}

func (r *employeeRepository) FindAll(limit, offset int, orderBy, sort string) ([]domain.Employee, int64, error) {
	var count int64
	err := r.db.Model(&domain.Employee{}).Count(&count).Error
	if err != nil {
		return nil, -1, err
	}

	queryOrder := fmt.Sprintf("%s %s", orderBy, sort)

	var employees []domain.Employee
	tx := r.db.Order(queryOrder).Limit(limit).Offset(offset).Find(&employees)
	if tx.Error != nil {
		return nil, -1, tx.Error
	}

	return employees, count, nil
}

func (r *employeeRepository) FindById(id uint) (domain.Employee, error) {
	var employee domain.Employee

	tx := r.db.Where("id", id).First(&employee)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return domain.Employee{}, ErrRecordNotFound
		}

		return domain.Employee{}, tx.Error
	}

	return employee, nil
}

func (r *employeeRepository) FindByEmail(email string) (domain.Employee, error) {
	var employee domain.Employee

	tx := r.db.Where("email", email).First(&employee)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return domain.Employee{}, ErrRecordNotFound
		}

		return domain.Employee{}, tx.Error
	}

	if employee.ID == 0 {
		return domain.Employee{}, ErrRecordNotFound
	}

	return employee, nil
}

func (r *employeeRepository) UpdateById(employee *domain.Employee) error {
	if employee == nil {
		return ErrNilReference
	}

	tx := r.db.Updates(employee)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (r *employeeRepository) DeleteById(id uint) error {
	var employee domain.Employee

	now := time.Now()
	tx := r.db.Model(&employee).Where("id", id).Updates(map[string]interface{}{
		"deleted_at": now,
	})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
