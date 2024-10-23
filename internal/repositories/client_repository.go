package repositories

import (
	"context"
	"database/sql"

	"BlueSoftBank/internal/models"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c *ClientRepository) GetClientByID(ctx context.Context, id int) (*models.Client, error) {
	query := `
		SELECT
	        id,
	        name,
	        email
		FROM 
		    clients 
		WHERE 
		    id = ?
    `
	row := c.db.QueryRowContext(ctx, query, id)

	client := &models.Client{}
	err := row.Scan(&client.ID, &client.Name, &client.Email)
	if err != nil {
		return nil, err
	}

	return client, nil
}
