package models

import "time"

type StaffDevelopmentPlan struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeId    int		`json:"employee_id,omitempty" form:"employee_id" validate:"omitempty,numeric"`
	FinancialYear    string		`json:"financial_year" form:"financial_year" validate:"required"`
	TrainingTitle    string		`json:"training_title" form:"training_title" validate:"required"`
	TrainingCategoryId    int		`json:"training_category_id,omitempty" form:"training_category_id" validate:"omitempty,numeric"`
	CountryId    int		`json:"country_id,omitempty" form:"country_id" validate:"omitempty,numeric"`
	InstitutionName    string		`json:"institution_name" form:"institution_name" validate:"required"`
	SponsorDetails    string		`json:"sponsor_details" form:"sponsor_details" validate:"required"`
	CourseDuration    string		`json:"course_duration" form:"course_duration" validate:"required"`
	CourseFee    float64		`json:"course_fee" form:"course_fee" validate:"required"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	PlanStatusId    int		`json:"plan_status_id,omitempty" form:"plan_status_id" validate:"omitempty,numeric"`
	SignificanceToInstitute    string		`json:"significance_to_institute" form:"significance_to_institute" validate:"required"`
	PlanFinalDecision    string		`json:"plan_final_decision" form:"plan_final_decision" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}