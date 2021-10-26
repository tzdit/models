package models

import "time"

//ActionType Structure
type ActionType struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name         string    `json:"name" form:"name" validate:"required"`
	Category     string    `json:"category" form:"category" validate:"required"`
	Descriptions string    `json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
