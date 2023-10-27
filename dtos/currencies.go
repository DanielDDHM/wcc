package dtos

type CurrencyDto struct {
	Name         string `json:"name" validate:"required,min=3,max=50"`
	IsoCode      string `json:"iso_code" validate:"required,min=3,max=6"`
	Type         string `json:"type" validate:"required" validate:"required,min=3,max=20"`
	Precision    int    `json:"precision" validate:"required,max=10"`
	Separator    string `json:"separator" validate:"required,max=10"`
	IsPivot      bool   `json:"is_pivot" validate:"required"`
	PriceInPivot string `json:"price_in_pivot" validate:"required"`
}
