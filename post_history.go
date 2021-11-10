package models

import "time"

type PostHistory struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId          int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	FromDate          time.Time `json:"from_date" form:"from_date" validate:"required"`
	ToDate            time.Time `json:"to_date" form:"to_date" validate:"required"`
	SubstantivePostId int       `json:"substantive_post_id,omitempty" form:"substantive_post_id" validate:"omitempty,numeric"`
	DutyPostId        int       `json:"duty_post_id,omitempty" form:"duty_post_id" validate:"omitempty,numeric"`
	CreatedBy         int       `json:"created_by,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

type PostHistoryFull struct {
	Id                   int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserTitle          string    `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	FromDate             time.Time `json:"from_date" form:"from_date" validate:"required"`
	ToDate               time.Time `json:"to_date" form:"to_date" validate:"required"`
	SubstantivePostTitle string    `json:"substantive_post_id,omitempty" form:"substantive_post_id" validate:"omitempty,numeric"`
	DutyPostTitle        string    `json:"duty_post_id,omitempty" form:"duty_post_id" validate:"omitempty,numeric"`
	CreatedBy            int       `json:"created_by,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
