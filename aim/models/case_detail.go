package models

import (
	"github.com/k0kubun/pp"
	"time"
)

//CaseDetail DataStructure
type CaseDetail struct {
	Id           int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Title        string    `json:"title" form:"title" validate:"required"`
	RefNo        string    `json:"ref_no" form:"ref_no" validate:"required"`
	From         string    `json:"from" form:"from" validate:"required"`
	Addressee    string    `json:"addressee" form:"addressee" validate:"required"`
	DateOnLetter string    `json:"date_on_letter,omitempty" form:"date_on_letter" validate:"required"`
	DateReceived string    `json:"date_received,omitempty" form:"date_received" validate:"required"`
	CreatedBy    int       `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type CustomTime time.Time

func (ct *CustomTime) UnmarshalParam(param string) error {
	pp.Printf("begin\n")
	t, err := time.Parse(`2006-01-02`, param)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	pp.Printf("end\n")
	return nil
}
