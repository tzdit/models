package models

import "time"

type FamilyJob struct{
	Id			int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	JobTitle    string		`json:"job_title" form:"job_title" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}