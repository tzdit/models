package models

import "time"

type MaritalStatus struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MaritalStatus    string		`json:"marital_status" form:"marital_status" validate:"required"`
	Description    string		`json:"description" form:"description" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}