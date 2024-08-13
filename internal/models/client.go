package models

// Client represents a bank client with multiple accounts
type Client struct {
	ID       string
	Name     string
	Accounts []*Account
}
