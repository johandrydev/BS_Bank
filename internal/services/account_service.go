package services

import (
	"BlueSoftBank/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type AccountService struct {
	db *sql.DB
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{db: db}
}

func (a *AccountService) CreateAccount(ctx context.Context, account *models.Account) error {
	return fmt.Errorf("method not implemented")
}
