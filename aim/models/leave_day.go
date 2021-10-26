package models

import "time"

type LeaveDay struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeRecordemployeeId    int		`json:"employee_recordemployee_id,omitempty" form:"employee_recordemployee_id" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	DateGranted    time.Time		`json:"date_granted" form:"date_granted" validate:"required"`
	DaysGranted    int		`json:"days_granted,omitempty" form:"days_granted" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}