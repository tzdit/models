package models

import "time"

type Node struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessListId    int		`json:"business_process_list_id,omitempty" form:"business_process_list_id" validate:"omitempty,numeric"`
	MyUserRoleId    int		`json:"my_user_role_id,omitempty" form:"my_user_role_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at"  validate:"required"`

}