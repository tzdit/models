package models

import "time"

//BusinessProcessInstance Structure
type BusinessProcessInstance struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyFileId              int       `json:"my_file_id,omitempty" form:"my_file_id" validate:"omitempty,numeric"`
	BusinessProcessListId int       `json:"business_process_list_id,omitempty" form:"business_process_list_id" validate:"omitempty,numeric"`
	MyUserId              int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	Timestamp             time.Time `json:"timestamp,omitempty" form:"timestamp"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}
