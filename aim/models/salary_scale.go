package models

import "time"

type SalaryScale struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ScaleCode    string    `json:"scale_code" form:"scale_code" validate:"required"`
	SalaryAmount float64   `json:"salary_amount" form:"salary_amount" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
