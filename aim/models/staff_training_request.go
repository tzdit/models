package models

import "time"

type StaffTrainingRequest struct{
	Id							int		`json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	StaffDevelopmentPlanId    	int		`json:"staff_development_plan_id,omitempty" form:"staff_development_plan_id" validate:"omitempty,numeric"`
	TrainingTitle    			string	`json:"training_title" form:"training_title" validate:"required"`
	TrainingCategoryId    		int		`json:"training_category_id,omitempty" form:"training_category_id" validate:"omitempty,numeric"`
	CountryId    				int		`json:"country_id,omitempty" form:"country_id" validate:"omitempty,numeric"`
	InstitutionName    			string		`json:"institution_name" form:"institution_name" validate:"required"`
	SponsorDetails    			string		`json:"sponsor_details" form:"sponsor_details" validate:"required"`
	CourseDuration    			float64		`json:"course_duration" form:"course_duration" validate:"required"`
	CourseFee    				float64		`json:"course_fee" form:"course_fee" validate:"required"`
	CurrencyId    				int		`json:"currency_id,omitempty" form:"currency_id" validate:"omitempty,numeric"`
	TrainStatusId    			int		`json:"train_status_id,omitempty" form:"train_status_id" validate:"omitempty,numeric"`
	SignificanceToInstitute    string		`json:"significance_to_institute" form:"significance_to_institute" validate:"required"`
	FromDate    				time.Time	`json:"from_date" form:"from_date" validate:"required"`
	ToDate    					time.Time	`json:"to_date" form:"to_date" validate:"required"`
	IsGranted    				bool		`json:"is_granted" form:"is_granted" validate:"required"`
	TrainingFinalDecision    	string		`json:"training_final_decision" form:"training_final_decision" validate:"required"`
	CreatedBy    				int		`json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	CreatedAt    				time.Time	`json:"created_at" form:"created_at" validate:"required"`
}
