package models

import "time"

type Profile struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId        int       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	PictureLocation string    `json:"picture_location" form:"picture_location" validate:"required"`
	SocialFb        string    `json:"social_fb" form:"social_fb" validate:"required"`
	SocialInstagram string    `json:"social_instagram" form:"social_instagram" validate:"required"`
	SocialTwitter   string    `json:"social_twitter" form:"social_twitter" validate:"required"`
	SocialLinkedin  string    `json:"social_linkedin" form:"social_linkedin" validate:"required"`
	CreatedBy       int       `json:"created_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}


type ProfileFull struct {
	Id              int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	MyUserId        string       `json:"my_user_id,omitempty" form:"my_user_id" validate:"omitempty,numeric"`
	PictureLocation string    `json:"picture_location" form:"picture_location" validate:"required"`
	SocialFb        string    `json:"social_fb" form:"social_fb" validate:"required"`
	SocialInstagram string    `json:"social_instagram" form:"social_instagram" validate:"required"`
	SocialTwitter   string    `json:"social_twitter" form:"social_twitter" validate:"required"`
	SocialLinkedin  string    `json:"social_linkedin" form:"social_linkedin" validate:"required"`
	CreatedBy       int       `json:"created_by,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}
