package models

import "time"

type Guideline struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	EmployeeRecordId    int		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	GuidelineTitle    string		`json:"guideline_title" form:"guideline_title" validate:"required"`
	GuidelineDescriptions    string		`json:"guideline_descriptions" form:"guideline_descriptions" validate:"required"`
	FinancialYear    string		`json:"financial_year" form:"financial_year" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`

}

type GuidelineFull struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeTitle   string		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	EmployeeRecordTitle   string		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	GuidelineTitle    string		`json:"guideline_title" form:"guideline_title" validate:"required"`
	GuidelineDescriptions    string		`json:"guideline_descriptions" form:"guideline_descriptions" validate:"required"`
	FinancialYear    string		`json:"financial_year" form:"financial_year" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
