package handlers

import (
	httpbs "BlueSoftBank/internal/pkg/http_bs"
	"BlueSoftBank/internal/services"
	"BlueSoftBank/internal/structures"
	"database/sql"
	"encoding/json"
	"net/http"
)

type movementHandler struct {
	movementSrv services.MovementCreator
}

func NewMovementHandler(db *sql.DB) *movementHandler {
	return &movementHandler{
		movementSrv: services.NewMovementService(db),
	}
}

func (m *movementHandler) CreateMovement(w http.ResponseWriter, r *http.Request) {
	var movementReq structures.MovementRequest
	err := json.NewDecoder(r.Body).Decode(&movementReq)
	if err := httpbs.ValidateErr(w, err, http.StatusBadRequest); err != nil {
		return
	}

	if !movementReq.IsValid() {
		httpbs.WriteError(w, http.StatusBadRequest, "Invalid movement request")
		return
	}

	movement := movementReq.ToMovement()
	// Call the movement service to create the movement
	err = m.movementSrv.CreateMovement(r.Context(), &movement)
	if err := httpbs.ValidateErr(w, err, http.StatusInternalServerError); err != nil {
		return
	}
	// Return the created movement
	httpbs.WriteJson(w, http.StatusCreated, movement, "Movement created succesfully")
}
