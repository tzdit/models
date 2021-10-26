package models

import "time"

type NewStaffRequirement struct{
	Id	int									`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	SubstantivePost    string				`json:"substantive_post" form:"substantive_post" validate:"required"`
	NoOfVaccancies    string				`json:"no_of_vaccancies" form:"no_of_vaccancies" validate:"required"`
	SalaryScaleId    int					`json:"salary_scale_id,omitempty" form:"salary_scale_id" validate:"omitempty,numeric"`
	Reasons    string						`json:"reasons" form:"reasons" validate:"required"`
	DepartmentId    int						`json:"department_id,omitempty" form:"department_id" validate:"omitempty,numeric"`
	AnnualSalary    float64					`json:"annual_salary" form:"annual_salary" validate:"required"`
	NewStaffRequirementStatusId    int		`json:"new_staff_requirement_status_id,omitempty" form:"new_staff_requirement_status_id" validate:"omitempty,numeric"`
	CreatedBy    int						`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time					`json:"created_at"  validate:"required"`

}
