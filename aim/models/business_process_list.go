package models

import (
	"time"
)

//BusinessProcessList Structure
type BusinessProcessList struct {
	Id         int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name       string    `json:"name" form:"name" validate:"required"`
	DutyPostId int       `json:"duty_post_id,omitempty" form:"duty_post_id" validate:"omitempty,numeric"`
	Purpose    string    `json:"purpose" form:"purpose" validate:"required"`
	Scope      string    `json:"scope" form:"scope" validate:"required"`
	CreatedBy  int       `json:"created_by,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}