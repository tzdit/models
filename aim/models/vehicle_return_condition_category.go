
package models

import "time"

type VehicleReturnConditionCategory struct{
	Id			int			`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Category    string		`json:"category" form:"category" validate:"required"`
	CreatedBy   int			`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt   time.Time	`json:"created_at"  validate:"required"`
}
