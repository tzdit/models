package models

import "time"

//DepartmentalObjective Structure
type DepartmentalObjective struct {
	Id                       int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	InstitutionalObjectiveId int       `json:"institutional_objective_id,omitempty" form:"institutional_objective_id" validate:"omitempty,numeric"`
	DepartmentId             int       `json:"department_id,omitempty" form:"department_id" validate:"omitempty,numeric"`
	Objective                string    `json:"objective" form:"objective" validate:"required"`
	Descriptions             string    `json:"descriptions" form:"descriptions" validate:"required"`
	StartDate                time.Time `json:"start_date,omitempty" form:"start_date" validate:"required"`
	EndDate                  time.Time `json:"end_date,omitempty" form:"end_date" validate:"required"`
	IsActive                 bool      `json:"is_active,omitempty" form:"is_active" validate:"required"`
	PercentCompleted         float64   `json:"percent_completed,omitempty" form:"percent_completed" validate:"required"`
	CreatedBy                int       `json:"created_by,omitempty"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
}

type DepartmentalObjectiveFull struct {
	Id                          int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	InstitutionalObjectiveTitle int       `json:"institutional_objective_id,omitempty" form:"institutional_objective_id" validate:"omitempty,numeric"`
	DepartmentTitle             int       `json:"department_id,omitempty" form:"department_id" validate:"omitempty,numeric"`
	Objective                   string    `json:"objective" form:"objective" validate:"required"`
	Descriptions                string    `json:"descriptions" form:"descriptions" validate:"required"`
	StartDate                   time.Time `json:"start_date,omitempty" form:"start_date" validate:"required"`
	EndDate                     time.Time `json:"end_date,omitempty" form:"end_date" validate:"required"`
	IsActive                    bool      `json:"is_active,omitempty" form:"is_active" validate:"required"`
	PercentCompleted            float64   `json:"percent_completed,omitempty" form:"percent_completed" validate:"required"`
	CreatedBy                   int       `json:"created_by,omitempty"`
	CreatedAt                   time.Time `json:"created_at,omitempty"`
}
