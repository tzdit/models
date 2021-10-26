package models

import (
	"time"
)

//DutyPost DataStructure
type DutyPost struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId int       `json:"action_type_id" form:"action_type_id" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
