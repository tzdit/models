package models

import "time"

type MyFile struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	FileName     string    `json:"file_name" form:"file_name" validate:"required"`
	FileNumber   string    `json:"file_number" form:"file_number" validate:"required"`
	MyFileTypeId int       `json:"my_file_type_id,omitempty" form:"my_file_type_id" validate:"omitempty,numeric"`
	FileStatusId int       `json:"file_status_id,omitempty" form:"file_status_id" validate:"omitempty,numeric"`
	StoredBy     int       `json:"stored_by,omitempty" form:"stored_by" validate:"omitempty,numeric"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
