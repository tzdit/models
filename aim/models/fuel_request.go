package models

import "time"

type FuelRequest struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleDetailvehicleId    int		`json:"vehicle_detailvehicle_id,omitempty" form:"vehicle_detailvehicle_id" validate:"omitempty,numeric"`
	EmployeeRecordemployeeId    int		`json:"employee_recordemployee_id,omitempty" form:"employee_recordemployee_id" validate:"omitempty,numeric"`
	VehicleUserId    int		`json:"vehicle_user_id,omitempty" form:"vehicle_user_id" validate:"omitempty,numeric"`
	DateRequested    time.Time		`json:"date_requested" form:"date_requested" validate:"required"`
	Reasons    string		`json:"reasons" form:"reasons" validate:"required"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Quantity    float64		`json:"quantity" form:"quantity" validate:"required"`
	TotalPrice    float64		`json:"total_price" form:"total_price" validate:"required"`
	MileageAtRequest    time.Time		`json:"mileage_at_request" form:"mileage_at_request" validate:"required"`
	FuelRequestStatusId    int		`json:"fuel_request_status_id,omitempty" form:"fuel_request_status_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}