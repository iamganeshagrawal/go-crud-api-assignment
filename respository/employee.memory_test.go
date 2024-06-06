package respository

import (
	"testing"

	"github.com/iamganeshagrawal/go-crud-api-assignment/models"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeInMemoryRepository_CreateEmployee(t *testing.T) {
	// Create a new in-memory repository
	repo := NewEmployeeInMemoryRepository()

	// Create an employee
	name, position, salary := "Ganesh Agrawal", "Software Engineer", 1234.00
	emp, err := repo.CreateEmployee(name, position, salary)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, 1, emp.ID, "ID should be 1")
	assert.Equal(t, name, emp.Name, "Name should be %s", name)
	assert.Equal(t, position, emp.Position, "Position should be %s", position)
	assert.Equal(t, salary, emp.Salary, "Salary should be %s", salary)

	// Create another employee
	name, position, salary = "Harshit Kumar", "DevOps Engineer", 1235.00
	emp, err = repo.CreateEmployee(name, position, salary)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, 2, emp.ID, "ID should be 2")
	assert.Equal(t, name, emp.Name, "Name should be %s", name)
	assert.Equal(t, position, emp.Position, "Position should be %s", position)
	assert.Equal(t, salary, emp.Salary, "Salary should be %s", salary)
}

func TestEmployeeInMemoryRepository_GetEmployeeByID(t *testing.T) {
	// Create a new in-memory repository
	repo := NewEmployeeInMemoryRepository()

	// Create an employee
	name, position, salary := "Ganesh Agrawal", "Software Engineer", 1234.00
	emp, _ := repo.CreateEmployee(name, position, salary)

	// Retrieve the employee by ID
	empByID, err := repo.GetEmployeeByID(emp.ID)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, emp.ID, empByID.ID, "ID should be %d", emp.ID)
	assert.Equal(t, emp.Name, empByID.Name, "Name should be %s", emp.Name)
	assert.Equal(t, emp.Position, empByID.Position, "Position should be %s", emp.Position)
	assert.Equal(t, emp.Salary, empByID.Salary, "Salary should be %s", emp.Salary)

	// Retrieve an employee by an invalid ID
	empByID, err = repo.GetEmployeeByID(2)

	assert.NotNil(t, err, "error should not be nil")
	assert.ErrorIs(
		t,
		err,
		ErrRecordNotFound,
		"error message should be 'record not found'",
	)
	assert.Equal(t, models.Employee{}, empByID, "employee should be empty")

	// Create another employee
	name, position, salary = "Harshit Kumar", "DevOps Engineer", 1235.00
	emp, _ = repo.CreateEmployee(name, position, salary)

	// Retrieve the employee by ID
	empByID, err = repo.GetEmployeeByID(emp.ID)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, emp.ID, empByID.ID, "ID should be %d", emp.ID)
	assert.Equal(t, emp.Name, empByID.Name, "Name should be %s", emp.Name)
	assert.Equal(t, emp.Position, empByID.Position, "Position should be %s", emp.Position)
	assert.Equal(t, emp.Salary, empByID.Salary, "Salary should be %s", emp.Salary)
}

func TestEmployeeInMemoryRepository_UpdateEmployee(t *testing.T) {
	// Create a new in-memory repository
	repo := NewEmployeeInMemoryRepository()

	// Create an employee
	name, position, salary := "Ganesh Agrawal", "Software Engineer", 1234.00
	emp, _ := repo.CreateEmployee(name, position, salary)

	// Update the employee
	newName, newPosition, newSalary := "Ganesh Agrawal", "Senior Software Engineer", 1350.00
	updatedEmp, err := repo.UpdateEmployee(emp.ID, newName, newPosition, newSalary)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, emp.ID, updatedEmp.ID, "ID should be %d", emp.ID)
	assert.Equal(t, newName, updatedEmp.Name, "Name should be %s", newName)
	assert.Equal(t, newPosition, updatedEmp.Position, "Position should be %s", newPosition)
	assert.Equal(t, newSalary, updatedEmp.Salary, "Salary should be %s", newSalary)

	// Fetch the updated employee
	empByID, err := repo.GetEmployeeByID(emp.ID)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, emp.ID, empByID.ID, "ID should be %d", emp.ID)
	assert.Equal(t, newName, empByID.Name, "Name should be %s", newName)
	assert.Equal(t, newPosition, empByID.Position, "Position should be %s", newPosition)
	assert.Equal(t, newSalary, empByID.Salary, "Salary should be %s", newSalary)

	// Update an employee with an invalid ID
	updatedEmp, err = repo.UpdateEmployee(2, newName, newPosition, newSalary)

	assert.NotNil(t, err, "error should not be nil")
	assert.ErrorIs(t, err, ErrRecordNotFound, "error message should be 'record not found'")
	assert.Equal(t, models.Employee{}, updatedEmp, "employee should be empty")
}

func TestEmployeeInMemoryRepository_DeleteEmployee(t *testing.T) {
	// Create a new in-memory repository
	repo := NewEmployeeInMemoryRepository()

	// Delete an employee with an invalid ID
	err := repo.DeleteEmployee(1)
	assert.NotNil(t, err, "error should not be nil")
	assert.ErrorIs(t, err, ErrRecordNotFound, "error message should be 'record not found'")

	// Create an employee
	name, position, salary := "Ganesh Agrawal", "Software Engineer", 1234.00
	emp, _ := repo.CreateEmployee(name, position, salary)

	// Delete the employee
	err = repo.DeleteEmployee(emp.ID)
	assert.Nil(t, err, "error should be nil")

	// Delete the employee again
	err = repo.DeleteEmployee(emp.ID)
	assert.NotNil(t, err, "error should not be nil")

	// Fetch the deleted employee
	empByID, err := repo.GetEmployeeByID(emp.ID)
	assert.NotNil(t, err, "error should not be nil")
	assert.ErrorIs(t, err, ErrRecordNotFound, "error message should be 'record not found'")
	assert.Equal(t, models.Employee{}, empByID, "employee should be empty")
}

func TestEmployeeInMemoryRepository_GetAllEmployees(t *testing.T) {
	// Create a new in-memory repository
	repo := NewEmployeeInMemoryRepository()

	// Fetch the all employees
	fetchedEmployees, total := repo.GetAllEmployees(0, 0)
	assert.Equal(t, 0, total, "total should be 0")
	assert.Empty(t, fetchedEmployees, "employees should be empty")

	// seed the repository with employees
	seedData := []models.Employee{
		{Name: "Ganesh Agrawal", Position: "Software Engineer", Salary: 1234.00},
		{Name: "Harshit Kumar", Position: "DevOps Engineer", Salary: 1235.00},
		{Name: "Rahul Singh", Position: "Data Scientist", Salary: 1236.00},
		{Name: "Rohit Sharma", Position: "Business Analyst", Salary: 1237.00},
		{Name: "Mahesh Kumar", Position: "QA Engineer", Salary: 1238.00},
		{Name: "Suresh Kumar", Position: "Technical Writer", Salary: 1239.00},
		{Name: "Ramesh Kumar", Position: "Network Engineer", Salary: 1240.00},
		{Name: "Rakesh Kumar", Position: "Security Analyst", Salary: 1241.00},
		{Name: "Pankaj Sharma", Position: "Software Engineer", Salary: 1242.00},
		{Name: "Ankit Sharma", Position: "Software Engineer", Salary: 1243.00},
		{Name: "Anil Sharma", Position: "Software Engineer", Salary: 1244.00},
	}

	for _, employee := range seedData {
		_, _ = repo.CreateEmployee(employee.Name, employee.Position, employee.Salary)
	}

	// Fetch the all employees
	fetchedEmployees, total = repo.GetAllEmployees(0, 0)
	assert.Equal(t, len(seedData), total, "total should be %d", len(seedData))
	assert.Equal(t, len(seedData), len(fetchedEmployees), "employees should be %d", len(seedData))

	// Fetch the first 5 employees
	fetchedEmployees, total = repo.GetAllEmployees(1, 5)
	assert.Equal(t, len(seedData), total, "total should be %d", len(seedData))
	assert.Equal(t, 5, len(fetchedEmployees), "employees should be 5")

	// Fetch the next 5 employees
	fetchedEmployees, total = repo.GetAllEmployees(2, 5)
	assert.Equal(t, len(seedData), total, "total should be %d", len(seedData))
	assert.Equal(t, 5, len(fetchedEmployees), "employees should be 5")
	for i, employee := range fetchedEmployees {
		assert.Equal(t, seedData[i+5].Name, employee.Name, "Name should be %s", seedData[i+5].Name)
		assert.Equal(
			t,
			seedData[i+5].Position,
			employee.Position,
			"Position should be %s",
			seedData[i+5].Position,
		)
		assert.Equal(
			t,
			seedData[i+5].Salary,
			employee.Salary,
			"Salary should be %s",
			seedData[i+5].Salary,
		)
	}

	// fetch the page that does not exist
	fetchedEmployees, total = repo.GetAllEmployees(99, 5)
	assert.Equal(t, len(seedData), total, "total should be %d", len(seedData))
	assert.Empty(t, fetchedEmployees, "employees should be empty")

	// fetch page with higher limit
	fetchedEmployees, total = repo.GetAllEmployees(1, 99)
	assert.Equal(t, len(seedData), total, "total should be %d", len(seedData))
	assert.Equal(t, len(seedData), len(fetchedEmployees), "employees should be %d", len(seedData))
}
