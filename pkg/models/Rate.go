package models

type Rate struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"Ccy"`
	BaseCurrency string  `json:"BaseCcy"`
}
