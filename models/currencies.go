package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CurrencyType string

type Currency struct {
	ID                uint             `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name" validate:"required,min=3,max=50"`
	IsoCode           string           `json:"iso_code" validate:"required,min=1,max=6"`
	Type              string           `json:"type" validate:"required,max=20"`
	Active            uint             `json:"active" gorm:"default:1"`
	Precision         int              `json:"precision" validate:"required"`
	Separator         string           `json:"separator" validate:"required"`
	IsPivot           uint             `json:"is_pivot"`
	PriceInPivot      string           `json:"price_in_pivot" validate:"required"`
	QuoteDestinations []*Quote         `json:"quote_destinations,omitempty" gorm:"ForeignKey:CurrencyDestinationId"`
	QuoteOrigins      []*Quote         `json:"quote_origins,omitempty" gorm:"ForeignKey:CurrencyOriginId"`
	CustomRatios      []*CustomRatio   `json:"custom_ratios,omitempty" gorm:"ForeignKey:CurrencyId"`
	QuoteCategories   []*QuoteCategory `json:"quote_categorias,omitempty" gorm:"ForeignKey:CurrencyId"`
	CreatedAt         *time.Time       `json:"created_at,omitempty"`
	UpdatedAt         *time.Time       `json:"updated_at,omitempty"`
}

func (currency *Currency) Prepare() error {
	err := currency.validate()

	if err != nil {
		return err
	}

	return nil
}

func (currency *Currency) validate() error {
	validate := validator.New()
	return validate.Struct(currency)
}
