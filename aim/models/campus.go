package models

import (
	"time"
)

//Campus DataStructure
type Campus struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	CampusTitle  string    `json:"campus_title" form:"campus_title" validate:"required"`
	CampusRegion string    `json:"campus_region" form:"campus_region" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
