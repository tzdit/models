package models

import "time"

type ExitRequestAction struct{
	Id					int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ExitRequestId    	int		`json:"exit_request_id,omitempty" form:"exit_request_id" validate:"omitempty,numeric"`
	PreviousId    		int		`json:"previous_id,omitempty" form:"previous_id" validate:"omitempty,numeric"`
	NextId    			int		`json:"next_id,omitempty" form:"next_id" validate:"omitempty,numeric"`
	ActionTypeId    	int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Minute    			string		`json:"minute" form:"minute" validate:"required"`
	CreatedBy    		int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    		time.Time		`json:"created_at" form:"created_at" validate:"required"`

}