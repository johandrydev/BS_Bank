package repositories

import (
	"database/sql"
	"fmt"
)

type MovementReposiroty struct {
	db *sql.DB
}

func (m *MovementReposiroty) CreateMovement() error {
	return fmt.Errorf("method not implemented")
}
