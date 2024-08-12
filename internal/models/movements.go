package models

type Movement struct {
	Type    string  `json:"type"`
	Amount  float64 `json:"amount"`
	Date    string  `json:"date"`
	City    string  `json:"city"`
	Account string  `json:"account"`
}

func (m *Movement) CreateMovement() error {
	return nil
}
