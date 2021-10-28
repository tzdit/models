package models

import "time"

type MyFileType struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	TypeName      string    `json:"type_name" form:"type_name" validate:"required"`
	Descriptions  string    `json:"descriptions" form:"descriptions" validate:"required"`
	KeywordFormat string    `json:"keyword_format" form:"keyword_format" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}
