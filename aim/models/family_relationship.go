package models

import "time"

type FamilyRelationship struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Relationship    string		`json:"relationship" form:"relationship" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}
