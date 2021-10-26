
package models

import "time"

type VehicleRequest struct{
	Id							int		`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	BusinessProcessInstanceId  	int		`json:"business_process_instance_id,omitempty" form:"business_process_instance_id" validate:"omitempty,numeric"`
	RequesterId    				int		`json:"requester_id,omitempty" form:"requester_id" validate:"omitempty,numeric"`
	VehicleRequestCategoryId    int		`json:"vehicle_request_category_id,omitempty" form:"vehicle_request_category_id" validate:"omitempty,numeric"`
	ReasonForRequest    		string		`json:"reason_for_request" form:"reason_for_request" validate:"required"`
	ContributionToInstitute    	string		`json:"contribution_to_institute" form:"contribution_to_institute" validate:"required"`
	FromDate    				time.Time	`json:"from_date" form:"from_date" validate:"required"`
	ToDate    					time.Time	`json:"to_date" form:"to_date" validate:"required"`
	FromMileage    				float64		`json:"from_mileage" form:"from_mileage" validate:"required"`
	ToMileage    				float64		`json:"to_mileage" form:"to_mileage" validate:"required"`
	FinalDecision    			string		`json:"final_decision" form:"final_decision" validate:"required"`
	VehicleId    				int		`json:"vehicle_id,omitempty" form:"vehicle_id" validate:"omitempty,numeric"`
	VehicleRequestStatusId    	int		`json:"vehicle_request_status_id,omitempty" form:"vehicle_request_status_id" validate:"omitempty,numeric"`
	AttachmentId    			int		`json:"attachment_id,omitempty" form:"attachment_id" validate:"omitempty,numeric"`
	CreatedBy    				int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    				time.Time	`json:"created_at" form:"created_at" validate:"required"`
}
