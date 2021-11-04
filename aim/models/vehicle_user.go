package models

import "time"

type VehicleUser struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleDetailId             int       `json:"vehicle_detail_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	Name                  string    `json:"name" form:"name" validate:"required"`
	Descriptions          string    `json:"descriptions" form:"descriptions" validate:"required"`
	FromDate              time.Time `json:"from_date" form:"from_date" validate:"required"`
	ToDate                time.Time `json:"to_date" form:"to_date" validate:"required"`
	VehicleUserAssignedBy int       `json:"vehicle_user_assigned_by,omitempty" form:"vehicle_user_assigned_by" validate:"omitempty,numeric"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}

//VehicleUserFull DataStructure
type VehicleUserFull struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleTitle          string    `json:"vehicle_detail_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	Name                  string    `json:"name" form:"name" validate:"required"`
	Descriptions          string    `json:"descriptions" form:"descriptions" validate:"required"`
	FromDate              time.Time `json:"from_date" form:"from_date" validate:"required"`
	ToDate                time.Time `json:"to_date" form:"to_date" validate:"required"`
	VehicleUserAssignedBy int       `json:"vehicle_user_assigned_by,omitempty" form:"vehicle_user_assigned_by" validate:"omitempty,numeric"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}
