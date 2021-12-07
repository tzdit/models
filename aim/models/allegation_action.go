package models

import "time"

//AllegationAction Structure
type AllegationAction struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name              string    `json:"name" form:"name" validate:"required"`
	Descriptions      string    `json:"descriptions" form:"descriptions" validate:"required"`
	AllegationLevelId int       `json:"allegation_level_id,omitempty" form:"allegation_level_id" validate:"omitempty,numeric"`
	AllegationId      int       `json:"allegation_id,omitempty" form:"allegation_id" validate:"omitempty,numeric"`
	CreatedBy         int       `json:"created_by,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

type AllegationActionFull struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name              string    `json:"name" form:"name" validate:"required"`
	Descriptions      string    `json:"descriptions" form:"descriptions" validate:"required"`
	AllegationLevelName string       `json:"allegation_level_id,omitempty" form:"allegation_level_id" validate:"omitempty,numeric"`
	AllegationId      string       `json:"allegation_id,omitempty" form:"allegation_id" validate:"omitempty,numeric"`
	CreatedBy         int       `json:"created_by,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}
