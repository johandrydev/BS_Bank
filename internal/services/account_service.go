package services

import (
	"BlueSoftBank/internal/models"
	"BlueSoftBank/internal/repositories"
	"context"
	"database/sql"
)

type AccountCreator interface {
	CreateAccount(ctx context.Context, account *models.Account) error
}

type AccountService struct {
	accountRepo AccountCreator
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{
		accountRepo: repositories.NewAccountRepository(db),
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, account *models.Account) error {
	return a.accountRepo.CreateAccount(ctx, account)
}
