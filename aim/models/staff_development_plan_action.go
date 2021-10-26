package models

import "time"

type StaffDevelopmentPlanAction struct {
	Id                     int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	StaffDevelopmentPlanId int       `json:"staff_development_plan_id,omitempty" form:"staff_development_plan_id" validate:"omitempty,numeric"`
	PreviousId             int       `json:"previous_id,omitempty" form:"previous_id" validate:"omitempty,numeric"`
	NextId                 int       `json:"next_id,omitempty" form:"next_id" validate:"omitempty,numeric"`
	ActionTypeId           int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Minute                 string    `json:"minute" form:"minute" validate:"required"`
	CreatedBy              int       `json:"created_by,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
}
