package respository

import (
	"github.com/iamganeshagrawal/go-crud-api-assignment/models"
)

// IEmployeeRepository is an interface for employee repository
type IEmployeeRepository interface {
	CreateEmployee(name string, position string, salary float64) (models.Employee, error)
	GetEmployeeByID(id int) (models.Employee, error)
	UpdateEmployee(id int, name string, position string, salary float64) (models.Employee, error)
	DeleteEmployee(id int) error
	GetAllEmployees(page int, limit int) ([]models.Employee, int)
}
