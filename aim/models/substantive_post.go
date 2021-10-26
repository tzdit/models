package models

import "time"

type SubstantivePost struct {
	Id                               int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	SubstantivePostTitle             string    `json:"substantive_post_title" form:"substantive_post_title" validate:"required"`
	SubstantivePostDescriptions      string    `json:"substantive_post_descriptions" form:"substantive_post_descriptions" validate:"required"`
	SubstantivePostEntryRequirements string    `json:"substantive_post_entry_requirements" form:"substantive_post_entry_requirements" validate:"required"`
	CreatedBy                        int       `json:"created_by,omitempty"`
	CreatedAt                        time.Time `json:"created_at"`
}
