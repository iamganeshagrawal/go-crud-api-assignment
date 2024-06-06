package main

import validation "github.com/go-ozzo/ozzo-validation/v4"

// CreateEmployeeRequest is the request body for creating an employee
type CreateEmployeeRequest struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func (form CreateEmployeeRequest) Validate() error {
	return validation.ValidateStruct(
		&form,
		validation.Field(&form.Name, validation.Required, validation.Length(3, 0)),
		validation.Field(&form.Position, validation.Required, validation.Length(3, 0)),
		validation.Field(&form.Salary, validation.Required, validation.Min(float64(0))),
	)
}

// UpdateEmployeeRequest is the request body for updating an employee
//
// It is the same as CreateEmployeeRequest. This is because the fields that can be updated.
// If you want to add more fields to the update request, you can do so here
type UpdateEmployeeRequest = CreateEmployeeRequest
