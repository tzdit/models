package models

import (
	"time"
)

//BusinessProcessInstanceTracking Structure
type BusinessProcessInstanceTracking struct {
	Id                        int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId int       `json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	Action                    string    `json:"action" form:"action" validate:"required"`
	Minutes                   string    `json:"minutes" form:"minutes" validate:"required"`
	MyUserId                  int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	Timestamp                 time.Time `json:"timestamp,omitempty" form:"timestamp" validate:"required"`
	CreatedBy                 int       `json:"created_by,omitempty"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	NodeId                    int       `json:"node_id,omitempty" form:"node_id" validate:"omitempty,numeric"`
}
