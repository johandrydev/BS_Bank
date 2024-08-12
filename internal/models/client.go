package models

// Client represents a bank client with multiple accounts
type Client struct {
	ID       string
	Name     string
	Accounts []*Account
}

// AddAccount adds an account to the client
func (c *Client) AddAccount(account *Account) {
	c.Accounts = append(c.Accounts, account)
}
