package models

import (
	"time"
)

//CaseDetail DataStructure
type CaseDetail struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Title        string    `json:"title" form:"title" validate:"required"`
	RefNo        string    `json:"ref_no" form:"ref_no" validate:"required"`
	From         string    `json:"from" form:"from" validate:"required"`
	Addressee    string    `json:"addressee" form:"addressee" validate:"required"`
	DateOnLetter time.Time `json:"date_on_letter,omitempty" form:"date_on_letter" validate:"required"`
	DateReceived time.Time `json:"date_received,omitempty" form:"date_received" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
