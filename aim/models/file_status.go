package models

import "time"

type FileStatus struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Status    string		`json:"status" form:"status" validate:"required"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}
