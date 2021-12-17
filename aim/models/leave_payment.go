package models

import "time"

type LeavePayment struct{
	Id							int		`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId   int		`json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	LeaveDetailId    			int		`json:"leave_detail_id,omitempty" form:"leave_detail_id" validate:"omitempty,numeric"`
	LeavePaymentItemId    		int		`json:"leave_payment_item_id,omitempty" form:"leave_payment_item_id" validate:"omitempty,numeric"`
	Quantity    				float64		`json:"quantity" form:"quantity" validate:"required"`
	TotalPrice    				float64		`json:"total_price" form:"total_price" validate:"required"`
	CreatedBy    				int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    				time.Time	`json:"created_at" form:"created_at" validate:"required"`
}

type LeavePaymentFull struct{
	Id							int		`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId   string		`json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	LeaveDetailId    			string		`json:"leave_detail_id,omitempty" form:"leave_detail_id" validate:"omitempty,numeric"`
	LeavePaymentItemId    		string		`json:"leave_payment_item_id,omitempty" form:"leave_payment_item_id" validate:"omitempty,numeric"`
	Quantity    				string		`json:"quantity" form:"quantity" validate:"required"`
	TotalPrice    				string		`json:"total_price" form:"total_price" validate:"required"`
	CreatedBy    				int			`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    				time.Time	`json:"created_at" form:"created_at" validate:"required"`
}