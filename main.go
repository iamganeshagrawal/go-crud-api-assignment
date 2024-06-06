package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/iamganeshagrawal/go-crud-api-assignment/respository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new in-memory repository
	empRepo := respository.NewEmployeeInMemoryRepository()

	// Create a new employee controller
	empController := NewEmployeeController(empRepo)

	// Create a new echo application
	app := echo.New()

	// add middleware
	app.Pre(middleware.RemoveTrailingSlash()) // Remove trailing slash from the URL
	app.Use(middleware.Logger())              // Log all requests
	app.Use(middleware.Recover())             // Recover from panics

	// Define routes
	// Grouping routes under /api/v1
	apiV1Group := app.Group("/api/v1")
	empGroup := apiV1Group.Group("/employees")

	// Define employee routes
	empGroup.PUT("/:id", empController.UpdateEmployee).Name = "employee.update"
	empGroup.DELETE("/:id", empController.DeleteEmployee).Name = "employee.delete"
	empGroup.GET("/:id", empController.GetEmployeeByID).Name = "employee.get"
	empGroup.POST("", empController.CreateEmployee).Name = "employee.create"
	empGroup.GET("", empController.GetAllEmployees).Name = "employee.list"

	// Ping or Health check endpoint
	app.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	}).Name = "ping"

	// List all routes in the application (For debugging)
	app.GET("/", func(c echo.Context) error {
		routes := app.Routes()
		return c.JSON(http.StatusOK, routes)
	}).Name = "index"

	// Start the echo application
	app.Logger.Fatal(app.Start(":8080"))
}

// EmployeeController is the controller for handling employee requests
type EmployeeController struct {
	repo respository.IEmployeeRepository
}

// NewEmployeeController creates a new employee controller
func NewEmployeeController(repo respository.IEmployeeRepository) *EmployeeController {
	return &EmployeeController{repo: repo}
}

// CreateEmployee creates a new employee
//
// POST /api/v1/employees
func (ec *EmployeeController) CreateEmployee(c echo.Context) error {
	var body CreateEmployeeRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request body",
		})
	}

	// Validate the request body
	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err})
	}

	// Create an employee from the request body
	emp, err := ec.repo.CreateEmployee(body.Name, body.Position, body.Salary)
	if err != nil {
		return err
	}

	// Return the created employee
	return c.JSON(http.StatusCreated, emp)
}

// GetEmployeeByID retrieves an employee by ID
//
// GET /api/v1/employees/:id
func (ec *EmployeeController) GetEmployeeByID(c echo.Context) error {
	// Get the employee ID from the URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid employee ID"})
	}

	// Retrieve the employee from the repository
	employee, err := ec.repo.GetEmployeeByID(id)
	if err != nil && errors.Is(err, respository.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "employee not found"})
	}
	if err != nil {
		return err
	}

	// Return the employee
	return c.JSON(http.StatusOK, employee)
}

// UpdateEmployee updates an employee by ID
//
// PUT /api/v1/employees/:id
func (ec *EmployeeController) UpdateEmployee(c echo.Context) error {
	// Get the employee ID from the URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid employee ID"})
	}

	var body UpdateEmployeeRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request body",
		})
	}

	// Validate the request body
	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err})
	}

	// Update the employee in the repository
	employee, err := ec.repo.UpdateEmployee(id, body.Name, body.Position, body.Salary)
	if err != nil && errors.Is(err, respository.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "employee not found"})
	}
	if err != nil {
		return err
	}

	// Return the updated employee
	// We can also return 204 No Content if we don't want to return the updated employee
	return c.JSON(http.StatusOK, employee)
}

// DeleteEmployee deletes an employee by ID
//
// DELETE /api/v1/employees/:id
func (ec *EmployeeController) DeleteEmployee(c echo.Context) error {
	// Get the employee ID from the URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid employee ID"})
	}

	// Delete the employee from the repository
	err = ec.repo.DeleteEmployee(id)
	if err != nil && errors.Is(err, respository.ErrRecordNotFound) {
		// Return 404 if the employee is not found
		// Or we can treat this as a successful deletion as well
		// and return 204 instead of 404
		return c.JSON(http.StatusNotFound, map[string]string{"error": "employee not found"})
	}
	if err != nil {
		return err
	}

	// Return 204 if the employee is successfully deleted
	return c.NoContent(http.StatusNoContent)
}

// GetAllEmployees retrieves all employees
//
// GET /api/v1/employees
func (ec *EmployeeController) GetAllEmployees(c echo.Context) error {
	// Get the page and limit query parameters
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	// Set default page to 1
	if pageStr == "" {
		pageStr = "1"
	}

	// Set default limit to 10
	if limitStr == "" {
		limitStr = "10"
	}

	// Convert page and limit to integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid page number"})
	}
	if page <= 0 {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "page number should be greater than 0"},
		)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid limit number"})
	}
	if limit <= 0 {
		limit = -1 // -1 means no limit
	}

	// Retrieve employees from the repository
	employees, total := ec.repo.GetAllEmployees(page, limit)

	// Create a list response
	response := ListResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Data:  employees,
	}

	// Return the list response
	return c.JSON(http.StatusOK, response)
}
