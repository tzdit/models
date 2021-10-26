package models

import "time"

type StaffAllegation struct {
	Id                        int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId int       `json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	InitatorId                int       `json:"initator_id,omitempty" form:"initator_id" validate:"omitempty,numeric"`
	AccussedId                int       `json:"accussed_id,omitempty" form:"accussed_id" validate:"omitempty,numeric"`
	AllegationId              int       `json:"allegation_id,omitempty" form:"allegation_id" validate:"omitempty,numeric"`
	IsActive                  bool      `json:"is_active" form:"is_active" validate:"required"`
	FinalDecision             string    `json:"final_decision" form:"final_decision" validate:"required"`
	AttachmentId              int       `json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy                 int       `json:"created_by,omitempty"`
	CreatedAt                 time.Time `json:"created_at"`
}
