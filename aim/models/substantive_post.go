package models

//SubstantivePost DataStructure
type SubstantivePost struct {
	SubstantivePostId                int    `json:"substantivepostid,omitempty" form:"substantivepostid" validate:"omitempty,numeric"`
	SubstantivePostTitle             string `json:"substantiveposttitle" form:"substantiveposttitle" validate:"required"`
	SubstantivePostDescriptions      string `json:"substantivepostdescriptions" form:"substantivepostdescriptions" validate:"required"`
	SubstantivePostEntryRequirements string `json:"substantivepostentryrequirements" form:"substantivepostentryrequirements" validate:"required"`
}
