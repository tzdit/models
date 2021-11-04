package models

import "time"

type FuelPrice struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTypeId  int       `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	PricePerLitre float64   `json:"price_per_litre" form:"price_per_litre" validate:"required"`
	Currency      string    `json:"currency" form:"currency" validate:"required"`
	FuelStatus    string    `json:"fuel_status" form:"fuel_status" validate:"required"`
	DateActivated time.Time `json:"date_activated" form:"date_activated" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type FuelPriceFull struct {
	Id            int       `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	ActionTitle   string    `json:"action_type_id,omitempty" form:"action_type_id" validate:"omitempty,numeric"`
	PricePerLitre float64   `json:"price_per_litre" form:"price_per_litre" validate:"required"`
	Currency      string    `json:"currency" form:"currency" validate:"required"`
	FuelStatus    string    `json:"fuel_status" form:"fuel_status" validate:"required"`
	DateActivated time.Time `json:"date_activated" form:"date_activated" validate:"required"`
	CreatedBy     int       `json:"created_by,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
