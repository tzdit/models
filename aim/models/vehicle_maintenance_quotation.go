package models

import "time"

type VehicleMaintenanceQuotation struct{
	Id						int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleInspectionId    	int		`json:"vehicle_inspection_id,omitempty" form:"vehicle_inspection_id" validate:"omitempty,numeric"`
	VendorId    			int		`json:"vendor_id,omitempty" form:"vendor_id" validate:"omitempty,numeric"`
	TotalPrice    			float64		`json:"total_price" form:"total_price" validate:"required"`
	IsApproved    			bool		`json:"is_approved" form:"is_approved" validate:"required"`
	CreatedBy    			int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    			time.Time		`json:"created_at" form:"created_at" validate:"required"`
}

type VehicleMaintenanceQuotationFull struct {
	Id						int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleInspectionId    	int		`json:"vehicle_inspection_id,omitempty" form:"vehicle_inspection_id" validate:"omitempty,numeric"`
	VendorId    			int		`json:"vendor_id,omitempty" form:"vendor_id" validate:"omitempty,numeric"`
	TotalPrice    			string		`json:"total_price" form:"total_price" validate:"required"`
	IsApproved    			string		`json:"is_approved" form:"is_approved" validate:"required"`
	CreatedBy    			int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    			time.Time		`json:"created_at" form:"created_at" validate:"required"`
}
