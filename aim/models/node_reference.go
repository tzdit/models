package models

import "time"

type NodeReference struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	NodeId    int		`json:"node_id,omitempty" form:"node_id" validate:"omitempty,numeric"`
	NextNodeId    int		`json:"next_node_id,omitempty" form:"next_node_id" validate:"omitempty,numeric"`
	PreviousNodeId    int		`json:"previous_node_id,omitempty" form:"previous_node_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at"  validate:"required"`
	BusinessProcessInstanceTrackingId    int		`json:"business_process_instance_tracking_id,omitempty" form:"business_process_instance_tracking_id" validate:"omitempty,numeric"`

}