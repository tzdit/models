
package models

import "time"

type VehicleMaintenanceRequest struct{
	Id					int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleId    		int			`json:"vehicle_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	DateRequested    	time.Time	`json:"date_requested" form:"date_requested" validate:"required"`
	ActionTypeId    	int			`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	MaintananceReasons  string		`json:"maintanance_reasons" form:"maintanance_reasons" validate:"required"`
	VehicleStatusId    	int			`json:"vehicle_status_id,omitempty" form:"vehicle_status_id" validate:"omitempty,numeric"`
	AttachmentId    	int			`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    		int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time	`json:"created_at" form:"created_at" validate:"required"`
}

type VehicleMaintenanceRequestFull struct {
	Id					int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleId    		string			`json:"vehicle_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	DateRequested    	time.Time	`json:"date_requested" form:"date_requested" validate:"required"`
	ActionTypeId    	string			`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	MaintananceReasons  string		`json:"maintanance_reasons" form:"maintanance_reasons" validate:"required"`
	VehicleStatusId    	string			`json:"vehicle_status_id,omitempty" form:"vehicle_status_id" validate:"omitempty,numeric"`
	AttachmentId    	string			`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    		int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time	`json:"created_at" form:"created_at" validate:"required"`
}
