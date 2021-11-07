package models

import "time"

type FuelRequest struct {
	Id                  int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleUserId       int       `json:"vehicle_user_id,omitempty" form:"vehicle_user_id" validate:"omitempty,numeric"`
	DateRequested       time.Time `json:"date_requested" form:"date_requested" validate:"required"`
	Reasons             string    `json:"reasons" form:"reasons" validate:"required"`
	ActionTypeId        int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Quantity            float64   `json:"quantity" form:"quantity" validate:"required"`
	TotalPrice          float64   `json:"total_price" form:"total_price" validate:"required"`
	MileageAtRequest    time.Time `json:"mileage_at_request" form:"mileage_at_request" validate:"required"`
	FuelRequestStatusId int       `json:"fuel_request_status_id,omitempty" form:"fuel_request_status_id" validate:"omitempty,numeric"`
	EmployeeRecordId    int       `json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	VehicleDetailId     int       `json:"vehicle_detail_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	CreatedBy           int       `json:"created_by,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
}

type FuelRequestFull struct {
	Id                     int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleUserTitle       string    `json:"vehicle_user_id,omitempty" form:"vehicle_user_id" validate:"omitempty,numeric"`
	DateRequested          time.Time `json:"date_requested" form:"date_requested" validate:"required"`
	Reasons                string    `json:"reasons" form:"reasons" validate:"required"`
	ActionTypeTitle        string    `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Quantity               float64   `json:"quantity" form:"quantity" validate:"required"`
	TotalPrice             float64   `json:"total_price" form:"total_price" validate:"required"`
	MileageAtRequest       time.Time `json:"mileage_at_request" form:"mileage_at_request" validate:"required"`
	FuelRequestStatusTitle string    `json:"fuel_request_status_id,omitempty" form:"fuel_request_status_id" validate:"omitempty,numeric"`
	EmployeeRecordTitle    string    `json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	VehicleDetailTitle     string    `json:"vehicle_detail_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	CreatedBy              int       `json:"created_by,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
}
