package models

import "time"

//EmployeeRecord DataStructure
type EmployeeRecord struct {
	Id                int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	PFNo              string    `json:"pf_no" form:"pf_no" validate:"required"`
	FirstName         string    `json:"first_name" form:"first_name" validate:"required"`
	MiddleName        string    `json:"middle_name" form:"middle_name" validate:"required"`
	LastName          string    `json:"last_name" form:"last_name" validate:"required"`
	SubstantivePostId int       `json:"substantive_post_id" form:"substantive_post_id" validate:"required"`
	DutyPostId        int       `json:"duty_post_id" form:"duty_post_id" validate:"required"`
	DepartmentId      int       `json:"department_id" form:"department_id" validate:"required"`
	DoB               time.Time `json:"dob" form:"dob" validate:"required"`
	PlaceOfBirth      string    `json:"place_of_birth" form:"place_of_birth" validate:"required"`
	DateEmployed      time.Time `json:"date_employed" form:"date_employed" validate:"required"`
	DateToRetire      time.Time `json:"date_to_retire" form:"date_to_retire" validate:"required"`
	PostalAddress     string    `json:"postal_address" form:"postal_address" validate:"required"`
	PhysicalAddress   string    `json:"physical_address" form:"physical_address" validate:"required"`
	PlaceOfDomicile   string    `json:"place_of_domicile" form:"place_of_domicile" validate:"required"`
	MaritalStatusId   int       `json:"marital_status_id" form:"marital_status_id" validate:"required"`
	ReligionId        int       `json:"religion_id" form:"religion_id" validate:"required"`
	ActionTypeId      int       `json:"action_type_id" form:"action_type_id" validate:"required"`
	SalaryScaleId     int       `json:"salary_scale_id" form:"salary_scale_id" validate:"required"`
	MyUserId          int       `json:"my_user_id" form:"my_user_id" validate:"required"`
	AppointmentDate   time.Time `json:"appointment_date" form:"appointment_date" validate:"required"`
	ConfirmationDate  time.Time `json:"confirmation_date" form:"confirmation_date" validate:"required"`
	CreatedBy         int       `json:"created_by,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
}

type EmployeeRecordFull struct {
	Id                   int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	PFNo                 string    `json:"pf_no" form:"pf_no" validate:"required"`
	FirstName            string    `json:"first_name" form:"first_name" validate:"required"`
	MiddleName           string    `json:"middle_name" form:"middle_name" validate:"required"`
	LastName             string    `json:"last_name" form:"last_name" validate:"required"`
	SubstantivePostTitle string    `json:"substantive_post_title" form:"substantive_post_id" validate:"required"`
	DutyPostTitle        string    `json:"duty_post_title" form:"duty_post_id" validate:"required"`
	DepartmentTitle      string    `json:"department_title" form:"department_id" validate:"required"`
	DoB                  time.Time `json:"dob" form:"dob" validate:"required"`
	PlaceOfBirth         string    `json:"place_of_birth" form:"place_of_birth" validate:"required"`
	DateEmployed         time.Time `json:"date_employed" form:"date_employed" validate:"required"`
	DateToRetire         time.Time `json:"date_to_retire" form:"date_to_retire" validate:"required"`
	PostalAddress        string    `json:"postal_address" form:"postal_address" validate:"required"`
	PhysicalAddress      string    `json:"physical_address" form:"physical_address" validate:"required"`
	PlaceOfDomicile      string    `json:"place_of_domicile" form:"place_of_domicile" validate:"required"`
	MaritalStatus        string    `json:"marital_status" form:"marital_status_id" validate:"required"`
	Religion             string    `json:"religion" form:"religion_id" validate:"required"`
	EmploymentType       string    `json:"employment_type" form:"action_type_id" validate:"required"`
	ScaleCode            string    `json:"scale_code" form:"salary_scale_id" validate:"required"`
	MyUserId             int       `json:"my_user_id" form:"my_user_id" validate:"required"`
	AppointmentDate      time.Time `json:"appointment_date" form:"appointment_date" validate:"required"`
	ConfirmationDate     time.Time `json:"confirmation_date" form:"confirmation_date" validate:"required"`
	CreatedBy            int       `json:"created_by,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
