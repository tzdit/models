package models

import "time"

type FamilyRelationship struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Relationship string    `json:"relationship" form:"relationship" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
