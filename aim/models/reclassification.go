package models

import "time"

type Reclassification struct {
	Id                            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId                  int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	SubstantivePostId             int       `json:"substantive_post_id,omitempty" form:"substantive_post_id" validate:"omitempty,numeric"`
	AttachmentId                  int       `json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	ReclassificationStatusId      int       `json:"reclassification_status_id,omitempty" form:"reclassification_status_id" validate:"omitempty,numeric"`
	ReclassificationFinalDecision string    `json:"reclassification_final_decision" form:"reclassification_final_decision" validate:"required"`
	CreatedBy                     int       `json:"created_by,omitempty"`
	CreatedAt                     time.Time `json:"created_at"`
}
