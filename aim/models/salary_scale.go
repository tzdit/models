package models

//SalaryScale DataStructure
type SalaryScale struct {
	SalaryScaleId int    `json:"salaryscaleid,omitempty" form:"salaryscaleid" validate:"omitempty,numeric"`
	ScaleCode     string `json:"scalecode" form:"scalecode" validate:"required"`
	SalaryAmount  string `json:"salaryamount" form:"salaryamount" validate:"required"`
}
