package repositories

import (
	"BlueSoftBank/internal/models"
	"context"
	"database/sql"
	"time"
)

type MovementReposiroty struct {
	db *sql.DB
}

func NewMovementRepository(db *sql.DB) *MovementReposiroty {
	return &MovementReposiroty{db: db}
}

func (m *MovementReposiroty) CreateMovement(ctx context.Context, movement *models.Movement) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(
		ctx,
		`INSERT INTO movements (
			account_id,
			type,
			amount,
			date,
			city
		) VALUES (
			$1, $2, $3
		)`,
		movement.AccountID,
		movement.Type,
		movement.Amount,
		time.Now(),
		movement.City,
	)
	return err
}
