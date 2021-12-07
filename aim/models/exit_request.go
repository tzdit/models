package models

import "time"

type ExitRequest struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId    int		`json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	TargetId    int		`json:"target_id,omitempty" form:"target_id" validate:"omitempty,numeric"`
	ActionTypeId    int		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	AttachmentId    int		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	IsFinished    bool		`json:"is_finished" form:"is_finished" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}

type ExitRequestfull struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId    int		`json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	TargetName    string		`json:"target_id,omitempty" form:"target_id" validate:"omitempty,numeric"`
	ActionTypeName    string		`json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	AttachmentName    string		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	IsFinished    bool		`json:"is_finished" form:"is_finished" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}