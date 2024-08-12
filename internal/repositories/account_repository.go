package repositories

import (
	"BlueSoftBank/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type AccountRepository struct {
	db *sql.DB
}

// Deposit adds funds to the account
func (a *AccountRepository) Deposit(ctx context.Context, account *models.Account, amount float64) error {
	// Implement the logic to deposit funds into the account
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := a.db.ExecContext(
		ctx,
		"UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, account.ID,
	)
	if err != nil {
		log.Println("Error trying to deposit in account")
		return fmt.Errorf("error trying to deposit in account: %v", err)
	}
	account.Balance += amount
	return nil
}

// Withdraw removes funds from the account
func (a *AccountRepository) Withdraw(ctx context.Context, account *models.Account, amount float64) error {
	// Implement the logic to withdraw funds from the account
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := a.db.ExecContext(
		ctx,
		"UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, account.ID,
	)
	if err != nil {
		log.Println("Error trying to withdraw from account")
		return fmt.Errorf("error trying to withdraw from account: %v", err)
	}
	account.Balance -= amount
	return nil
}

// CheckBalance returns the current balance of the account
func (a *AccountRepository) CheckBalance(ctx context.Context, number string) *models.Account {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Implement the logic to retrieve the account balance
	var account models.Account
	_ = a.db.QueryRowContext(
		ctx,
		"SELECT id, balance FROM accounts WHERE number = ?", number,
	).Scan(&account.ID, &account.Balance)
	// Update the account balance in the database
	return &account
}
