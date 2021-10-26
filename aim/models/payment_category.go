package models
import "time"

type PaymentCategory struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Category    string		`json:"category" form:"category" validate:"required"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" validate:"required"`

}