package services

import (
	"context"
	"database/sql"

	"BlueSoftBank/internal/models"
	"BlueSoftBank/internal/repositories"
)

type ClientService struct {
	ClientRepo *repositories.ClientRepository
}

func NewClientService(db *sql.DB) *ClientService {
	return &ClientService{
		ClientRepo: repositories.NewClientRepository(db),
	}
}

func (c *ClientService) GetClientByID(ctx context.Context, id int) (*models.Client, error) {
	return c.ClientRepo.GetClientByID(ctx, id)
}
