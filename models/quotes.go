package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Quote struct {
	ID                    uint           `json:"id" gorm:"primaryKey"`
	PartnerId             uint           `json:"partnerId" gorm:"index,not null"`
	UserIdentifier        string         `json:"user_identifier" gorm:"not null"`
	Partner               *Partner       `json:"partner,omitempty"`
	QuoteId               string         `json:"quote_id" gorm:"index,not null"`
	Price                 string         `json:"price" gorm:"type:varchar(255);not null"`
	ExpiresIn             time.Time      `json:"expires_in" gorm:"not null"`
	Status                string         `json:"status" gorm:"default:pending;not null"`
	CurrencyDestinationId uint           `json:"currency_destination_id" gorm:"not null"`
	CurrencyDestination   *Currency      `json:"currency_destination,omitempty"`
	CurrencyOriginId      uint           `json:"currency_origin_id" gorm:"not null"`
	CurrencyOrigin        *Currency      `json:"currency_origin,omitempty"`
	Size                  string         `json:"size" gorm:"type:varchar(50);inot null"`
	PartnerResult         datatypes.JSON `json:"partner_result " gorm:"not null"`
	CreatedAt             *time.Time     `json:"created_at" gorm:"not null"`
	UpdatedAt             *time.Time     `json:"updated_at" gorm:"not null"`
}

func (quote *Quote) Prepare() error {
	quote.QuoteId = uuid.New().String()

	err := quote.validate()
	if err != nil {
		return err
	}
	return nil
}

func (quote *Quote) validate() error {
	validate := validator.New()
	return validate.Struct(quote)
}
