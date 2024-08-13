package services

import (
	"BlueSoftBank/internal/models"
	"BlueSoftBank/internal/repositories"
	"context"
	"database/sql"
	"log"
)

type MovementCreator interface {
	CreateMovement(ctx context.Context, movement *models.Movement) error
}

type AccountUpdater interface {
	Deposit(ctx context.Context, accountId int, amount float64) error
	Withdraw(ctx context.Context, accountId int, amount float64) error
}

type movementService struct {
	movementRepo MovementCreator
	AccountRepo  AccountUpdater
}

func NewMovementService(db *sql.DB) *movementService {
	return &movementService{
		movementRepo: repositories.NewMovementRepository(db),
		AccountRepo:  repositories.NewAccountRepository(db),
	}
}

func (s *movementService) CreateMovement(ctx context.Context, movement *models.Movement) error {
	// create movement with a transaction and after create a deposit or windraw in the account repository
	err := s.movementRepo.CreateMovement(ctx, movement)
	if err != nil {
		log.Println(err)
		return err
	}

	if movement.Type == "deposit" {
		err := s.AccountRepo.Deposit(ctx, movement.AccountID, movement.Amount)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	// withdraw
	err = s.AccountRepo.Withdraw(ctx, movement.AccountID, movement.Amount)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
