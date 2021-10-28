package models

import "time"

type DutyPost struct {
	Id                  int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	CreatedBy           int       `json:"created_by,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	DutyPostTitle       string    `json:"duty_post_title" form:"duty_post_title" validate:"required"`
	DutyPostDescription string    `json:"duty_post_description" form:"duty_post_description" validate:"required"`
}
