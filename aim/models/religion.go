package models

import "time"

type Religion struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Religion     string    `json:"religion" form:"religion" validate:"required"`
	Descriptions string    `json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
