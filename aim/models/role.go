package models

import (
	"aim/services/entity"
)

//Role DataStructure
type Role struct {
	ID          entity.ID `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description"`
}

type RolePermissions struct {
	ID          entity.ID `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Permissions []int32   `json:"permissions" form:"permissions[]" validate:"required"`
}
