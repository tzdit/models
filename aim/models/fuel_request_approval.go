package models

import "time"

type FuelRequestApproval struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	FuelRequestId    int		`json:"fuel_request_id,omitempty" form:"fuel_request_id" validate:"omitempty,numeric"`
	ApprovingOfficer    string		`json:"approving_officer" form:"approving_officer" validate:"required"`
	FuelRequestActionId    int		`json:"fuel_request_action_id,omitempty" form:"fuel_request_action_id" validate:"omitempty,numeric"`
	Minutes    string		`json:"minutes" form:"minutes" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}

type FuelRequestApprovalFull struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	FuelRequestTitle    string		`json:"fuel_request_id,omitempty" form:"fuel_request_id" validate:"omitempty,numeric"`
	ApprovingOfficer    string		`json:"approving_officer" form:"approving_officer" validate:"required"`
	FuelRequestActionTitle    string		`json:"fuel_request_action_id,omitempty" form:"fuel_request_action_id" validate:"omitempty,numeric"`
	Minutes    string		`json:"minutes" form:"minutes" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}