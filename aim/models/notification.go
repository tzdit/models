package models

import "time"

type Notification struct {
	Id                     int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId               int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	TargetUserId           int       `json:"target_user_id,omitempty" form:"target_user_id" validate:"omitempty,numeric"`
	NotificationCategoryId int       `json:"notification_category_id,omitempty" form:"notification_category_id" validate:"omitempty,numeric"`
	NotificationTitle      string    `json:"notification_title" form:"notification_title" validate:"required"`
	Descriptions           string    `json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy              int       `json:"created_by,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
}

type NotificationFull struct {
	Id                       int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId                 int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	TargetUserId             int       `json:"target_user_id,omitempty" form:"target_user_id" validate:"omitempty,numeric"`
	NotificationCategoryName string    `json:"notification_category_id,omitempty" form:"notification_category_id" validate:"omitempty,numeric"`
	NotificationTitle        string    `json:"notification_title" form:"notification_title" validate:"required"`
	Descriptions             string    `json:"descriptions" form:"descriptions" validate:"required"`
	CreatedBy                int       `json:"created_by,omitempty"`
	CreatedAt                time.Time `json:"created_at"`
}
