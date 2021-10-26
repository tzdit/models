package models

import "time"

type PayrollApproval struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	PayrollId    int		`json:"payroll_id,omitempty" form:"payroll_id" validate:"omitempty,numeric"`
	ApprovingOfficer    string		`json:"approving_officer" form:"approving_officer" validate:"required"`
	PayrollActionId    int		`json:"payroll_action_id,omitempty" form:"payroll_action_id" validate:"omitempty,numeric"`
	Minutes    string		`json:"minutes" form:"minutes" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" validate:"required"`

}
