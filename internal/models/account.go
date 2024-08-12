package models

// Account represents a bank account
type Account struct {
	ID        int
	Number    string
	Client    Client
	Balance   float64
	movements []Movement
}
