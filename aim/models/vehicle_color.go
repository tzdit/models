package models

import "time"

type VehicleColor struct{
	Id			int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ColorName   string		`json:"color_name" form:"color_name" validate:"required"`
	CreatedBy   int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt   time.Time	`json:"created_at" form:"created_at" validate:"required"`
}