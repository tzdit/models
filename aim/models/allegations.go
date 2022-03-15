package models

import (
	"time"
)

//Allegation Structure
type Allegation struct {
	Id           int       `json:"id" form:"id" validate:"omitempty,numeric"`
	ActionTypeId int       `json:"action_type_id" form:"action_type_id" validate:"omitempty,numeric"`
	DateAccused  time.Time `json:"date_accused" form:"date_accused"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type AllegationFull struct {
	Id           int       `json:"id" form:"id" validate:"omitempty,numeric"`
	ActionTypeId string       `json:"action_type_id" form:"action_type_id" validate:"omitempty,numeric"`
	DateAccused  time.Time `json:"date_accused" form:"date_accused"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
