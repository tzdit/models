package models

import "time"

type VehicleStatus struct{
	Id				int	`		json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name    		string		`json:"name" form:"name" validate:"required"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    	int			`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    	time.Time	`json:"created_at"  validate:"required"`
}