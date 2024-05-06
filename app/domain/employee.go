package domain

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID        uint         		 `gorm:"column:id;autoIncrement;primaryKey"`
	FirstName string       		 `gorm:"column:first_name"`
	LastName  string       		 `gorm:"column:last_name"`
	Email     string       		 `gorm:"column:email;index"`
	HireDate  time.Time   		 `gorm:"column:hire_date;type:date;index"`
	CreatedAt *time.Time   		 `gorm:"column:created_at"`
	UpdatedAt *time.Time   		 `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt    `gorm:"column:deleted_at;index"`
}

type EmployeeRequest struct {
	FirstName string       		 `json:"first_name"`
	LastName  string       		 `json:"last_name"`
	Email     string       		 `json:"email"`
	HireDate  string 		     `json:"hire_date"`
}

type EmployeeResponse struct {
	Id        uint  			 `json:"id"`
	FirstName string       		 `json:"first_name"`
	LastName  string       		 `json:"last_name"`
	Email     string       		 `json:"email"`
	HireDate  time.Time   		 `json:"hire_date"`
	CreatedAt *time.Time   		 `json:"created_at,omitempty"`
	UpdatedAt *time.Time   		 `json:"updated_at,omitempty"`
}

type PaginationResponse struct {
	PageNum 	int			`json:"page_number"`
	PageSize 	int			`json:"page_size"`
	TotalPage	int64		`json:"total_page"`	
	Data 		interface{}	`json:"data"`
}
