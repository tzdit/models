package models

import "time"

//DriverVehicle Structure
type DriverVehicle struct {
	Id               int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleDetailId  int       `json:"vehicle_detail_id,omitempty" form:"vehicle_detail_id" validate:"omitempty,numeric"`
	EmployeeRecordId int       `json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	DateFrom         time.Time `json:"date_from" form:"date_from" validate:"required"`
	DateTo           time.Time `json:"date_to" form:"date_to" validate:"required"`
	DriverAssignedBy int       `json:"driver_assigned_by,omitempty" form:"driver_assigned_by" validate:"omitempty,numeric"`
	CreatedBy        int       `json:"created_by,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}
