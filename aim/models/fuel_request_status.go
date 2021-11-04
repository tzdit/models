package models

import "time"

type FuelRequestStatus struct {
	Id        int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Status    string    `json:"status" form:"status" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
