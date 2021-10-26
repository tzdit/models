package models

import "time"

type VehicleDetail struct{
	Id				int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name    		string		`json:"name" form:"name" validate:"required"`
	Model    		string		`json:"model" form:"model" validate:"required"`
	Year    		int			`json:"year,omitempty" form:"year" validate:"omitempty,numeric"`
	VehicleBrandId  int			`json:"vehicle_brand_id,omitempty" form:"vehicle_brand_id" validate:"omitempty,numeric"`
	VehicleColorId  int			`json:"vehicle_color_id,omitempty" form:"vehicle_color_id" validate:"omitempty,numeric"`
	RegistrationId  int			`json:"registration_id,omitempty" form:"registration_id" validate:"omitempty,numeric"`
	VehicleTypeId   int			`json:"vehicle_type_id,omitempty" form:"vehicle_type_id" validate:"omitempty,numeric"`
	CampusId    	int			`json:"campus_id,omitempty" form:"campus_id" validate:"omitempty,numeric"`
	VehicleStatusId int			`json:"vehicle_status_id,omitempty" form:"vehicle_status_id" validate:"omitempty,numeric"`
	CampusAssignedBy    int		`json:"campus_assigned_by,omitempty" form:"campus_assigned_by" validate:"omitempty,numeric"`
	Mileage    			float64		`json:"mileage" form:"mileage" validate:"required"`
	AttachmentId    	int		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	EmployeeId    		int		`json:"employee_id,omitempty" form:"employee_id" validate:"omitempty,numeric"`
	CreatedBy    		int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time		`json:"created_at" form:"created_at" validate:"required"`
}
