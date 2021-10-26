package models

import (
	"time"
)

//EmployeeObjective DataStructure
type EmployeeObjective struct {
	Id               int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeRecordId int       `json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	Objective        string    `json:"objective" form:"objective" validate:"required"`
	Descriptions     string    `json:"descriptions" form:"descriptions" validate:"required"`
	StartDate        time.Time `json:"start_date,omitempty" form:"start_date" validate:"required"`
	EndDate          time.Time `json:"end_date,omitempty" form:"end_date" validate:"required"`
	IsActive         bool      `json:"is_active" form:"is_active" validate:"required"`
	IsRevised        bool      `json:"is_revised" form:"is_revised" validate:"required"`
	AssignedBy       int       `json:"assigned_by,omitempty" form:"assigned_by" validate:"omitempty,numeric"`
	PercentCompleted int       `json:"percent_completed,omitempty" form:"percent_completed" validate:"omitempty,numeric"`
	CreatedBy        int       `json:"created_by,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}