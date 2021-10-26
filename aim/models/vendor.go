package models

import "time"

type Vendor struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name            string    `json:"name" form:"name" validate:"required"`
	Address         string    `json:"address" form:"address" validate:"required"`
	Tin             string    `json:"tin" form:"tin" validate:"required"`
	PhoneNumber     string    `json:"phone_number" form:"phone_number" validate:"required"`
	Email           string    `json:"email" form:"email" validate:"required"`
	Website         string    `json:"website" form:"website" validate:"required"`
	PhysicalAddress string    `json:"physical_address" form:"physical_address" validate:"required"`
	ActionTypeId    int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	CreatedBy       int       `json:"created_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}
