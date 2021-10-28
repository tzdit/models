package models

import "time"

type MaritalStatus struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MaritalStatus string    `json:"marital_status" form:"marital_status" validate:"required"`
	Description   string    `json:"description" form:"description" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
