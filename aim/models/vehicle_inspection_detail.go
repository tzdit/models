package models

import "time"

type VehicleInspectionDetail struct{
	Id						int		`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleInspectionId    	int		`json:"vehicle_inspection_id,omitempty" form:"vehicle_inspection_id" validate:"omitempty,numeric"`
	ProblemObserved    		string	`json:"problem_observed" form:"problem_observed" validate:"required"`
	Descriptions    		string	`json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    			int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    			time.Time	`json:"created_at" form:"created_at" validate:"required"`
}
