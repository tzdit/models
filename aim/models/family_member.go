package models

import "time"

type FamilyMember struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId    int		`json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	FullName    string		`json:"full_name" form:"full_name" validate:"required"`
	FamilyRelationshipId    int		`json:"family_relationship_id,omitempty" form:"family_relationship_id" validate:"omitempty,numeric"`
	FamilyJobId    int		`json:"family_job_id,omitempty" form:"family_job_id" validate:"omitempty,numeric"`
	Email    string		`json:"email" form:"email" validate:"required"`
	Phone    string		`json:"phone" form:"phone" validate:"required"`
	AltenatePhone    string		`json:"altenate_phone" form:"altenate_phone" validate:"required"`
	PostalAddress    string		`json:"postal_address" form:"postal_address" validate:"required"`
	PhysicalAddress    string		`json:"physical_address" form:"physical_address" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`
}