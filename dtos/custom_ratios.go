package dtos

type CustomRatioDto struct {
	Target    string  `json:"target" validate:"required"`
	Currency  string  `json:"currency" validate:"required"`
	AskRatio  float64 `json:"ask_ratio" validate:"required"`
	BuyRatio  float64 `json:"buy_ratio" validate:"required"`
	PartnerId uint    `json:"partner_id" validate:"required"`
}
