package models

import "time"

type ReclassificationStatus struct {
	Id                     int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ReclassificationStatus string    `json:"reclassification_status" form:"reclassification_status" validate:"required"`
	CreatedBy              int       `json:"created_by,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
}
