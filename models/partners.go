package models

import (
	"errors"
	"time"

	"github.com/DanielDDHM/world-coin-converter/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Partner struct {
	ID              uint             `json:"id" gorm:"primaryKey"`
	Name            string           `json:"name" validate:"required,min=3,max=50"`
	PartnerKey      string           `json:"partner_key,omitempty" validate:"required"`
	PartnerId       string           `json:"partner_id,omitempty" validate:"required"`
	Country         string           `json:"country" validate:"required,min=3,max=50"`
	Locale          string           `json:"locale" validate:"required,min=2,max=5"`
	Ttl             uint             `json:"ttl" validate:"required"`
	Active          uint             `json:"active" gorm:"default:1"`
	Quotes          []*Quote         `json:"quotes,omitempty" gorm:"ForeignKey:PartnerId"`
	QuoteCategories []*QuoteCategory `json:"quote_categories,omitempty" gorm:"ForeignKey:PartnerId"`
	CreatedAt       *time.Time       `json:"created_at"`
	UpdatedAt       *time.Time       `json:"updated_at"`
}

func (partner *Partner) Prepare() error {
	if partner.Ttl > 10 {
		return errors.New("Ttl must be less than 10")
	}
	partner.PartnerId = uuid.New().String()
	partner.PartnerKey = utils.RandomString(50)

	err := partner.validate()
	if err != nil {
		return err
	}
	return nil
}

func (partner *Partner) validate() error {
	validate := validator.New()
	return validate.Struct(partner)
}
