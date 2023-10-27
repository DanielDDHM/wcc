package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuoteCategory struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" validate:"required,min=3,max=50"`
	SpreadAsk  string     `json:"spread_ask" validate:"required"`
	SpreadBid  string     `json:"spread_bid" validate:"required"`
	Ttls       int        `json:"ttls" validate:"required" gorm:"not null"`
	CurrencyId uint       `json:"currency_id" validate:"required" gorm:"not null"`
	Currency   *Currency  `json:"currency,omitempty"`
	PartnerId  uint       `json:"partner_id" validate:"required"`
	Partner    *Partner   `json:"partner,omitempty"`
	CreatedAt  *time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"not null"`
}

func (quoteCategory *QuoteCategory) Prepare() error {
	err := quoteCategory.validate()

	if err != nil {
		return err
	}

	return nil
}

func (quoteCategory *QuoteCategory) validate() error {
	validate := validator.New()
	return validate.Struct(quoteCategory)
}
