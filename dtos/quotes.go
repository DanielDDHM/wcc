package dtos

type RequestQuoteDto struct {
	Amount              float64 `json:"amount" validate:"required"`
	OriginCurrency      string  `json:"origin_currency" validate:"required"`
	DestinationCurrency string  `json:"destination_currency" validate:"required"`
	Size                float64 `json:"size" validate:"required"`
	UserIdentifier      string  `json:"user_identifier"`
}

type RequestQuoteOwsDto struct {
	Trading    string  `json:"trading"`
	Settlement string  `json:"settlement"`
	Side       string  `json:"side"`
	Size       float64 `json:"size"`
	Amount     float64 `json:"amount"`
}
