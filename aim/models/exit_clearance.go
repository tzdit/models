package models

import "time"

type ExitClearance struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ExitRequestId     int       `json:"exit_request_id,omitempty" form:"exit_request_id" validate:"omitempty,numeric"`
	ActionTypeId      int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	ExitClearanceDone bool      `json:"exit_clearance_done" form:"exit_clearance_done" validate:"required"`
	CreatedBy         int       `json:"created_by,omitempty" validate:"omitempty,numeric"`
	CreatedAt         time.Time `json:"created_at"`
}
