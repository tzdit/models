package models

import "time"

type PhysicalFileTracking struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyFileId    int		`json:"my_file_id,omitempty" form:"my_file_id" validate:"omitempty,numeric"`
	To    string		`json:"to" form:"to" validate:"required"`
	ReceivedTimestamp    time.Time		`json:"received_timestamp" form:"received_timestamp" validate:"required"`
	DispatchedTimestamp    time.Time		`json:"dispatched_timestamp" form:"dispatched_timestamp" validate:"required"`
	FileRequestedBy    int		`json:"file_requested_by,omitempty" form:"file_requested_by" validate:"omitempty,numeric"`
	RequestTimestamp    time.Time		`json:"request_timestamp" form:"request_timestamp" validate:"required"`
	RequiredDate    time.Time		`json:"required_date" form:"required_date" validate:"required"`
	IsAttended    time.Time		`json:"is_attended" form:"is_attended" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at"  validate:"required"`

}