package structures

import "BlueSoftBank/internal/models"

type MovementRequest struct {
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	City      string  `json:"city"`
	AccountID int     `json:"account_id"`
}

func (m MovementRequest) IsValid() bool {
	if m.Type != "deposit" && m.Type != "withdraw" {
		return false
	}
	return m.Amount > 0 && m.City != "" && m.AccountID > 0
}

func (m MovementRequest) ToMovement() models.Movement {
	return models.Movement{
		Amount:    m.Amount,
		Type:      m.Type,
		City:      m.City,
		AccountID: m.AccountID,
	}
}
