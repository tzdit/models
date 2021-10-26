package models

import "time"

type LeaveDetail struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeRecordemployeeId    int		`json:"employee_recordemployee_id,omitempty" form:"employee_recordemployee_id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId    int		`json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	DateFrom    time.Time		`json:"date_from" form:"date_from" validate:"required"`
	DateTo    time.Time		`json:"date_to" form:"date_to" validate:"required"`
	IsActive    bool		`json:"is_active" form:"is_active" validate:"required"`
	AttachmentId    int		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	IsPaid    bool		`json:"is_paid" form:"is_paid" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}