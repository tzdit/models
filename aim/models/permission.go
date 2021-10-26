package models

import (
	"time"
)

// Permission Struct
type Permission struct {
	ID         int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Path       string    `json:"path" form:"path" validate:"required"`
	Method     string    `json:"method" form:"method" validate:"required"`
	Service    string    `json:"service" form:"service" validate:"required"`
	SubService string    `json:"sub_service" form:"sub_service" validate:"required"`
	Action     string    `json:"action" form:"action" validate:"required"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
