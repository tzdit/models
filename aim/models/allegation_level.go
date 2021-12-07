package models

import "time"

//AllegationLevel Structure
type AllegationLevel struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name         string    `json:"name" form:"name" validate:"required"`
	Descriptions string    `json:"descriptions" form:"descriptions" validate:"required"`
	MyUserRoleId int       `json:"my_user_role_id,omitempty" form:"my_user_role_id" validate:"omitempty,numeric"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
type AllegationLevelFull struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name         string    `json:"name" form:"name" validate:"required"`
	Descriptions string    `json:"descriptions" form:"descriptions" validate:"required"`
	MyUserRoleId int       `json:"my_user_role_id,omitempty" form:"my_user_role_id" validate:"omitempty,numeric"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}