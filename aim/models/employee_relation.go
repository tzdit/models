package models

import "time"

//EmployeeRelation DataStructure
type EmployeeRelation struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	PayrollId    int       `json:"payroll_id,omitempty" form:"payroll_id" validate:"omitempty,numeric"`
	MyUserId     int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	ActionTypeId int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
