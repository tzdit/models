package models

import (
	"aim/services/entity"
	"time"
)

// UserCategory DataStructure
type UserCategory struct {
	ID        entity.ID `json:"id,omitempty" form:"id" validate:"omitempty,required"`
	Name      string    `json:"name" form:"name" validate:"required"`
	CreatedBy int32     `json:"created_by" form:"created_by" validate:"required,numeric"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
