package models

import "time"

type VehicleInspection struct{
	Id				int						`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleMaintenanceRequestId    int		`json:"vehicle_maintenance_request_id,omitempty" form:"vehicle_maintenance_request_id" validate:"omitempty,numeric"`
	Inspector    	string					`json:"inspector" form:"inspector" validate:"required"`
	DateInspected  	time.Time				`json:"date_inspected" form:"date_inspected" validate:"required"`
	IsRepairable    bool					`json:"is_repairable" form:"is_repairable" validate:"required"`
	AttachmentId    int						`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    	int						`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    	time.Time				`json:"created_at" form:"created_at" validate:"required"`
}