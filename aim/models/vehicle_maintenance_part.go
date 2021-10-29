package models

import "time"

type VehicleMaintenancePart struct {
	Id                  int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleInspectionId int       `json:"vehicle_inspection_id,omitempty" form:"vehicle_inspection_id" validate:"omitempty,numeric"`
	PartName            string    `json:"part_name" form:"part_name" validate:"required"`
	PartDescriptions    string    `json:"part_descriptions" form:"part_descriptions" validate:"required"`
	Quantity            float64   `json:"quantity" form:"quantity" validate:"required"`
	UnitPrice           float64   `json:"unit_price" form:"unit_price" validate:"required"`
	VendorId            int       `json:"vendor_id,omitempty" form:"vendor_id" validate:"omitempty,numeric"`
	TotalPrice          float64   `json:"total_price" form:"total_price" validate:"required"`
	CreatedBy           int       `json:"created_by,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
}

type VehicleMaintenancePartFull struct {
	Id                  int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleInspectionId int       `json:"vehicle_inspection_id,omitempty" form:"vehicle_inspection_id" validate:"omitempty,numeric"`
	PartName            string    `json:"part_name" form:"part_name" validate:"required"`
	PartDescriptions    string    `json:"part_descriptions" form:"part_descriptions" validate:"required"`
	Quantity            float64   `json:"quantity" form:"quantity" validate:"required"`
	UnitPrice           float64   `json:"unit_price" form:"unit_price" validate:"required"`
	VendorName          string    `json:"vendor_id,omitempty" form:"vendor_id" validate:"omitempty,numeric"`
	TotalPrice          float64   `json:"total_price" form:"total_price" validate:"required"`
	CreatedBy           int       `json:"created_by,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
}

