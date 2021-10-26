package models

import "time"

//Case DataStructure
type Case struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyFileId     int       `json:"my_file_id,omitempty" form:"my_file_id" validate:"omitempty,numeric"`
	CaseTitle    string    `json:"case_title" form:"case_title" validate:"required"`
	ActionTypeId int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	CaseDetailId int       `json:"case_detail_id,omitempty" form:"case_detail_id" validate:"omitempty,numeric"`
	MyUserId     int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	CaseStatusId int       `json:"case_status_id,omitempty" form:"case_status_id" validate:"omitempty,numeric"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
