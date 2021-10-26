package models
import "time"

type PerformanceAssessmentRecommendation struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeRecordId    int		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	StartDate    time.Time		`json:"start_date" form:"start_date" validate:"required"`
	EndDate    time.Time		`json:"end_date" form:"end_date" validate:"required"`
	Recommendations    string		`json:"recommendations" form:"recommendations" validate:"required"`
	RecommendationBy    int		`json:"recommendation_by,omitempty" form:"recommendation_by" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at"  validate:"required"`

}