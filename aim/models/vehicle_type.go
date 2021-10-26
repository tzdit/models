package models

import "time"

type VehicleType struct{
	Id					int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name    			string		`json:"name" form:"name" validate:"required"`
	Descriptions    	string		`json:"descriptions" form:"descriptions" validate:"required"`
	Capacity    		int			`json:"capacity,omitempty" form:"capacity" validate:"omitempty,numeric"`
	CreatedBy    		int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time	`json:"created_at" form:"created_at" validate:"required"`

}

