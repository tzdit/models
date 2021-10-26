package models

import "time"

//Attachment Structure
type Attachment struct {
	Id                    int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId              int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	BusinessProcessListId int       `json:"business_process_list_id,omitempty" form:"business_process_list_id" validate:"omitempty,numeric"`
	ProcessInstanceId     int       `json:"process_instance_id,omitempty" form:"process_instance_id" validate:"omitempty,numeric"`
	ActionTypeId          int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	Title                 string    `json:"title" form:"title" validate:"required"`
	Path                  string    `json:"path" form:"path" validate:"required"`
	CreatedBy             int       `json:"created_by,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
}
