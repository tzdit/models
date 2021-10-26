package models

import "time"

type ReclassificationPlan struct {
	Id                       int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeRecordId         int       `json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	CurrentSubstantivePost   string    `json:"current_substantive_post" form:"current_substantive_post" validate:"required"`
	NewSubstantivePost       string    `json:"new_substantive_post" form:"new_substantive_post" validate:"required"`
	SalaryScaleId            int       `json:"salary_scale_id,omitempty" form:"salary_scale_id" validate:"omitempty,numeric"`
	ActionTypeId             int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	AnnualSalary             float64   `json:"annual_salary" form:"annual_salary" validate:"required"`
	Reasons                  string    `json:"reasons" form:"reasons" validate:"required"`
	ReclassificationStatusId int       `json:"reclassification_status_id,omitempty" form:"reclassification_status_id" validate:"omitempty,numeric"`
	CreatedBy                int       `json:"created_by,omitempty"`
	CreatedAt                time.Time `json:"created_at"`
}
