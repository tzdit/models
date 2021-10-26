package models


import "time"

type PerformanceAssessment struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	EmployeeObjectiveId    int		`json:"employee_objective_id,omitempty" form:"employee_objective_id" validate:"omitempty,numeric"`
	FactorsAffectingPerformance    string		`json:"factors_affecting_performance" form:"factors_affecting_performance" validate:"required"`
	SelfRating    string		`json:"self_rating" form:"self_rating" validate:"required"`
	SupervisorComments    string		`json:"supervisor_comments" form:"supervisor_comments" validate:"required"`
	SupervisorRating    string		`json:"supervisor_rating" form:"supervisor_rating" validate:"required"`
	AgreedRating    string		`json:"agreed_rating" form:"agreed_rating" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at"  validate:"required"`
	EmployeeRecordId    int		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`

}
