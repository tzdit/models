package models

import "time"

type VehicleRequestApproval struct{
	Id				int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	EmployeeId    	int			`json:"employee_id,omitempty" form:"employee_id" validate:"omitempty,numeric"`
	VehicleRequestActionId    int		`json:"vehicle_request_action_id,omitempty" form:"vehicle_request_action_id" validate:"omitempty,numeric"`
	Minutes    		string		`json:"minutes" form:"minutes" validate:"required"`
	VehicleRequestId    int		`json:"vehicle_request_id,omitempty" form:"vehicle_request_id" validate:"omitempty,numeric"`
	CreatedBy    	int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    	time.Time	`json:"created_at" form:"created_at" validate:"required"`
}