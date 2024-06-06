package respository

import (
	"fmt"
	"sync"

	"github.com/iamganeshagrawal/go-crud-api-assignment/internal/datatypes"
	"github.com/iamganeshagrawal/go-crud-api-assignment/models"
)

// Ensure type implements the interface
var _ IEmployeeRepository = (*EmployeeInMemoryRepository)(nil)

// EmployeeInMemoryRepository is an in-memory repository for employees
type EmployeeInMemoryRepository struct {
	mu     *sync.RWMutex                               // Mutex for thread-safety
	store  *datatypes.OrderedMap[int, models.Employee] // In-memory database
	nextId int                                         // Next available ID for the next employee
}

// NewEmployeeInMemoryRepository creates a new in-memory repository for employees
func NewEmployeeInMemoryRepository() *EmployeeInMemoryRepository {
	return &EmployeeInMemoryRepository{
		mu:     &sync.RWMutex{},
		store:  datatypes.NewOrderedMap[int, models.Employee](),
		nextId: 1,
	}
}

// CreateEmployee creates a new employee
func (repo *EmployeeInMemoryRepository) CreateEmployee(
	name string,
	position string,
	salary float64,
) (models.Employee, error) {
	repo.mu.Lock()         // Lock the mutex
	defer repo.mu.Unlock() // Unlock the mutex when the function returns

	// Create a new employee
	employee := models.Employee{
		ID:       repo.nextId, // Assign the next available ID
		Name:     name,
		Position: position,
		Salary:   salary,
	}

	// Store the employee in the store (in-memory database)
	repo.store.Set(employee.ID, employee)

	// Increment the next available ID
	repo.nextId++

	// Return the created employee
	return employee, nil
}

// GetEmployeeByID retrieves an employee by ID
func (repo *EmployeeInMemoryRepository) GetEmployeeByID(id int) (models.Employee, error) {
	repo.mu.RLock()         // Lock the mutex for reading
	defer repo.mu.RUnlock() // Unlock the mutex when the function returns

	// Retrieve the employee from the store
	employee, ok := repo.store.Get(id)
	if !ok {
		return models.Employee{}, fmt.Errorf(
			"employee with ID %d not found: %w",
			id,
			ErrRecordNotFound,
		)
	}

	// Return the employee
	return employee, nil
}

// UpdateEmployee updates an employee by ID
func (repo *EmployeeInMemoryRepository) UpdateEmployee(
	id int,
	name string,
	position string,
	salary float64,
) (models.Employee, error) {
	repo.mu.Lock()         // Lock the mutex
	defer repo.mu.Unlock() // Unlock the mutex when the function returns

	// Retrieve the employee from the store
	employee, ok := repo.store.Get(id)
	if !ok {
		return models.Employee{}, fmt.Errorf(
			"employee with ID %d update failed: %w",
			id,
			ErrRecordNotFound,
		)
	}

	// Update the employee
	employee.Name = name
	employee.Position = position
	employee.Salary = salary

	// Store the updated employee in the store
	repo.store.Set(employee.ID, employee)

	// Return the updated employee
	return employee, nil
}

// DeleteEmployee deletes an employee by ID
func (repo *EmployeeInMemoryRepository) DeleteEmployee(id int) error {
	repo.mu.Lock()         // Lock the mutex
	defer repo.mu.Unlock() // Unlock the mutex when the function returns

	// Delete the employee from the store
	if ok := repo.store.Delete(id); !ok {
		return fmt.Errorf(
			"employee with ID %d delete failed: %w",
			id,
			ErrRecordNotFound,
		)
	}

	// Return nil (no error)
	return nil
}

// GetAllEmployees retrieves all employees and total count
func (repo *EmployeeInMemoryRepository) GetAllEmployees(
	page int,
	limit int,
) ([]models.Employee, int) {
	repo.mu.RLock()         // Lock the mutex for reading
	defer repo.mu.RUnlock() // Unlock the mutex when the function returns

	// Handle pagination
	if page <= 0 {
		// If page is less than or equal to 0, set it to 1
		page = 1
	}
	if limit <= 0 {
		// If limit is less than or equal to 0, return all employees
		limit = repo.store.Len()
	}
	// Calculate the offset for pagination
	offset := (page - 1) * limit

	// Check if the offset is out of bounds
	if offset >= repo.store.Len() {
		return []models.Employee{}, repo.store.Len()
	}

	// Retrieve all employees from the store (in-memory database) with pagination
	employees := make([]models.Employee, 0, limit)

	empIds := repo.store.Keys()
	empIds = empIds[offset:]

	for i, id := range empIds {
		employee, _ := repo.store.Get(id)
		employees = append(employees, employee)
		if i == limit-1 {
			break
		}
	}

	// Return all employees
	return employees, repo.store.Len()
}
