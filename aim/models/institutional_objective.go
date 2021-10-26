package models

import "time"

type InstitutionalObjective struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Objective    string		`json:"objective" form:"objective" validate:"required"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	StartDate    time.Time		`json:"start_date" form:"start_date" validate:"required"`
	EndDate    time.Time		`json:"end_date" form:"end_date" validate:"required"`
	IsActive    bool		`json:"is_active" form:"is_active" validate:"required"`
	PercentCompleted    float64		`json:"percent_completed" form:"percent_completed" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" form:"created_at" validate:"required"`

}