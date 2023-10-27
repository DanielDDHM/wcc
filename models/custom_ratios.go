package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CustomRatio struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Target     string     `json:"target" validate:"required"`
	CurrencyId uint       `json:"currency_id" validate:"required"`
	Currency   *Currency  `json:"currency,omitempty"`
	PartnerId  uint       `json:"partner_id" validate:"required"`
	Partner    *Partner   `json:"partner,omitempty"`
	AskRatio   float64    `json:"ask_ratio" validate:"required"`
	BuyRatio   float64    `json:"buy_ratio" validate:"required"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (custom *CustomRatio) Prepare() error {
	err := custom.validate()
	if err != nil {
		return err
	}
	return nil
}

func (custom *CustomRatio) validate() error {
	validate := validator.New()
	return validate.Struct(custom)
}

// Só os funcionário do DanielDDHM

// target pode ser o CPF/CNPJ => Qualquer coisa que identifique o usuário
// Ask Ratio => Spread ask que está no quote categories
// Bid Ratio => Spread bid que está no quote categories
