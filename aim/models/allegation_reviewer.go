package models

import (
	"time"
)

//AllegationReviewer Structure
type AllegationReviewer struct {
	Id                     int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	StaffAllegationId      int       `json:"staff_allegation_id,omitempty" form:"staff_allegation_id" validate:"omitempty,numeric"`
	ReviewDate             time.Time `json:"review_date,omitempty" form:"review_date" validate:"required"`
	ReviewerId             int       `json:"reviewer_id,omitempty" form:"reviewer_id" validate:"omitempty,numeric"`
	AllegationActionId     int       `json:"allegation_action_id,omitempty" form:"allegation_action_id" validate:"omitempty,numeric"`
	ReviewerRecommendation string    `json:"reviewer_recommendation" form:"reviewer_recommendation" validate:"required"`
	CreatedBy              int       `json:"created_by,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
}