package structures

import "BlueSoftBank/internal/models"

type AccountRequest struct {
	ID      int     `json:"id"`
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}

func (ar *AccountRequest) ToAccount() models.Account {
	return models.Account{
		ID:      ar.ID,
		Number:  ar.Number,
		Balance: ar.Balance,
	}
}
