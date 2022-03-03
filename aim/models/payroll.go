package models

import "time"

type Payroll struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	CheckNo    int		`json:"check_no,omitempty" form:"check_no" validate:"omitempty,numeric"`
	EmployeeRecordId    int		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	PaymentCategoryId    int		`json:"payment_category_id,omitempty" form:"payment_category_id" validate:"omitempty,numeric"`
	Amount    float64		`json:"amount,omitempty" form:"amount" validate:"omitempty,numeric"`
	PreparedBy    int		`json:"prepared_by,omitempty" form:"prepared_by" validate:"omitempty,numeric"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	PayMonth    string		`json:"pay_month" form:"pay_month" validate:"required"`
	PayrollStatusId    int		`json:"payroll_status_id,omitempty" form:"payroll_status_id" validate:"omitempty,numeric"`
	AttachmentId    int		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" validate:"required"`

}


type PayrollFull struct{
	Id	int	`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	CheckNo    int		`json:"check_no,omitempty" form:"check_no" validate:"omitempty,numeric"`
	EmployeeRecordTitle    string		`json:"employee_record_id,omitempty" form:"employee_record_id" validate:"omitempty,numeric"`
	PaymentCategoryTitle   string		`json:"payment_category_id,omitempty" form:"payment_category_id" validate:"omitempty,numeric"`
	Amount    float64		`json:"amount,omitempty" form:"amount" validate:"omitempty,numeric"`
	PreparedBy    int		`json:"prepared_by,omitempty" form:"prepared_by" validate:"omitempty,numeric"`
	Descriptions    string		`json:"descriptions" form:"descriptions" validate:"required"`
	PayMonth    string		`json:"pay_month" form:"pay_month" validate:"required"`
	PayrollStatusTitle    string		`json:"payroll_status_id,omitempty" form:"payroll_status_id" validate:"omitempty,numeric"`
	AttachmentTitle    string		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    int		`json:"created_by,omitempty"  validate:"omitempty,numeric"`
	CreatedAt    time.Time		`json:"created_at" validate:"required"`

}