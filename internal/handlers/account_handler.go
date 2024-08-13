package handlers

import (
	httpbs "BlueSoftBank/internal/pkg/http_bs"
	"BlueSoftBank/internal/services"
	"BlueSoftBank/internal/structures"
	"database/sql"
	"encoding/json"
	"net/http"
)

type accountHandler struct {
	accountService services.AccountCreator
}

func NewAccountHandler(db *sql.DB) *accountHandler {
	return &accountHandler{
		accountService: services.NewAccountService(db),
	}
}

func (a *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountReq structures.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountReq)
	if err := httpbs.ValidateErr(w, err, http.StatusBadRequest); err != nil {
		return
	}

	account := accountReq.ToAccount()
	ctx := r.Context()
	err = a.accountService.CreateAccount(ctx, &account)
	if err := httpbs.ValidateErr(w, err, http.StatusBadRequest); err != nil {
		return
	}
	httpbs.WriteJson(w, http.StatusCreated, account, "Account created succesfully")
}
