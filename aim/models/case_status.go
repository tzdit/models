package models

import "time"

//CaseStatus DataStructure
type CaseStatus struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	StatusName   string    `json:"status_name" form:"status_name" validate:"required"`
	Descriptions string    `json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
