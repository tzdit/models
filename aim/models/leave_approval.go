package models

import "time"

type LeaveApproval struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	LeaveDetailId int       `json:"leave_detail_id,omitempty" form:"leave_detail_id" validate:"omitempty,numeric"`
	DutyPostId    int       `json:"duty_post_id,omitempty" form:"duty_post_id" validate:"omitempty,numeric"`
	ActionTypeId  int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Minutes       string    `json:"minutes" form:"minutes" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

type LeaveApprovalFull struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	LeaveDetailId int       `json:"leave_detail_id,omitempty" form:"leave_detail_id" validate:"omitempty,numeric"`
	DutyPostTitle string    `json:"duty_post_id,omitempty" form:"duty_post_id" validate:"required"`
	ActionName    string    `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Minutes       string    `json:"minutes" form:"minutes" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
