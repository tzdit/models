package models

import "time"

//ActionDetail Structure
type ActionDetail struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	CurrentUserId int       `json:"current_user_id,omitempty" form:"current_user_id" validate:"omitempty,numeric"`
	NextUserId    int       `json:"next_user_id,omitempty" form:"next_user_id" validate:"omitempty,numeric"`
	Minutes       string    `json:"minutes" form:"minutes" validate:"required"`
	AssignedTime  time.Time `json:"assigned_time,omitempty" form:"assigned_time" validate:"required"`
	DispatchTime  time.Time `json:"dispatch_time,omitempty" form:"dispatch_time" validate:"required"`
	ActionTypeId  int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	CaseId        int       `json:"case_id,omitempty" form:"case_id" validate:"omitempty,numeric"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
