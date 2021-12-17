package models

import (
	"time"
)

//Allegation Structure
type Allegation struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	DateAccused  time.Time `json:"date_accused,omitempty" form:"date_accused" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type AllegationFull struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId string       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	DateAccused  time.Time `json:"date_accused,omitempty" form:"date_accused" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
