
package models

import "time"

type VehicleReturnCondition struct{
	Id					int				`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	VehicleRequestId    int				`json:"vehicle_request_id,omitempty" form:"vehicle_request_id" validate:"omitempty,numeric"`
	VehicleReturnConditionCategoryId    int		`json:"vehicle_return_condition_category_id,omitempty" form:"vehicle_return_condition_category_id" validate:"omitempty,numeric"`
	Remarks    			string			`json:"remarks" form:"remarks" validate:"required"`
	CreatedBy    		int				`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time		`json:"created_at" form:"created_at" validate:"required"`
}
