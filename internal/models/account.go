package models

// Account represents a bank account
type Account struct {
	ID       int
	Number   string
	Balance  float64
	ClientID int
}
