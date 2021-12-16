package models

import "time"

type VehicleDetail struct {
	Id               int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name             string    `json:"name" form:"name" validate:"required"`
	Model            string    `json:"model" form:"model" validate:"required"`
	Year             int       `json:"year,omitempty" form:"year" validate:"omitempty,numeric"`
	VehicleBrandId   int       `json:"vehicle_brand_id,omitempty" form:"vehicle_brand_id" validate:"omitempty,numeric"`
	VehicleColorId   int       `json:"vehicle_color_id,omitempty" form:"vehicle_color_id" validate:"omitempty,numeric"`
	RegistrationId   string    `json:"registration_id,omitempty" form:"registration_id" validate:"required"`
	VehicleTypeId    int       `json:"vehicle_type_id,omitempty" form:"vehicle_type_id" validate:"omitempty,numeric"`
	CampusId         int       `json:"campus_id,omitempty" form:"campus_id" validate:"omitempty,numeric"`
	VehicleStatusId  int       `json:"vehicle_status_id,omitempty" form:"vehicle_status_id" validate:"omitempty,numeric"`
	CampusAssignedBy string       `json:"campus_assigned_by,omitempty" form:"campus_assigned_by" validate:"omitempty,numeric"`
	Mileage          float64   `json:"mileage" form:"mileage" validate:"required"`
	AttachmentId     int       `json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	EmployeeId       int       `json:"employee_id,omitempty" form:"employee_id" validate:"omitempty,numeric"`
	CreatedBy        int       `json:"created_by,omitempty"`
	CreatedAt        time.Time `json:"created_at" `
}

type VehicleDetailFull struct {
	Id               int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name             string    `json:"name" form:"name" validate:"required"`
	Model            string    `json:"model" form:"model" validate:"required"`
	Year             int       `json:"year,omitempty" form:"year" validate:"omitempty,numeric"`
	BrandName        string    `json:"vehicle_brand_id,omitempty" form:"vehicle_brand_id" validate:"required"`
	ColorName        string    `json:"vehicle_color_id,omitempty" form:"vehicle_color_id" validate:"required"`
	RegistrationId   string    `json:"registration_id,omitempty" form:"registration_id" validate:"required"`
	VehicleType      string    `json:"vehicle_type_id,omitempty" form:"vehicle_type_id" validate:"required"`
	CampusTitle      string    `json:"campus_id,omitempty" form:"campus_id" validate:"required"`
	VehicleStatus    string    `json:"vehicle_status_id,omitempty" form:"vehicle_status_id" validate:"required"`
	CampusAssignedBy string    `json:"campus_assigned_by,omitempty" form:"campus_assigned_by" validate:"required"`
	Mileage          float64   `json:"mileage" form:"mileage" validate:"required"`
	AttachmentTitle  string    `json:"attachment_id,omitempty" form:"attachment_id" validate:"required"`
	EmployeeName     string    `json:"employee_id,omitempty" form:"employee_id" validate:"required"`
	CreatedBy        int       `json:"created_by,omitempty"`
	CreatedAt        time.Time `json:"created_at" `
}
