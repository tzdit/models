package models

import "time"

//Country Structure
type Country struct {
	Id        int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name      string    `json:"name" form:"name" validate:"required"`
	Continent string    `json:"continent" form:"continent" validate:"required"`
	CreatedBy int       `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
