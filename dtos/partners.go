package dtos

type CreatePartnerDto struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Locale  string `json:"locale"`
	Ttl     uint   `json:"ttl"`
}
