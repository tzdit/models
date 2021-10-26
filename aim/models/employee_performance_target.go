package models

import "time"

//EmployeePerformanceTarget DataStructure
type EmployeePerformanceTarget struct {
	Id                  int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeObjectiveId int       `json:"employee_objective_id" form:"employee_objective_id" validate:"required"`
	ExpectedOutput      string    `json:"expected_output" form:"expected_output" validate:"required"`
	Target              string    `json:"target" form:"target" validate:"required"`
	Indicator           string    `json:"indicator" form:"indicator" validate:"required"`
	Resources           string    `json:"resources" form:"resources" validate:"required"`
	CreatedBy           int       `json:"created_by,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
}
