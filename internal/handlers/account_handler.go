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
	db             *sql.DB
	accountService services.AccountService
}

func NewAccountHandler(db *sql.DB) *accountHandler {
	return &accountHandler{
		db:             db,
		accountService: *services.NewAccountService(db),
	}
}

func (a *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountReq structures.AccountRequest
	if err := httpbs.ValidateErr(w, json.NewDecoder(r.Body).Decode(&accountReq), http.StatusBadRequest); err != nil {
		return
	}
	account := accountReq.ToAccount()
	ctx := r.Context()
	if err := httpbs.ValidateErr(w,
		a.accountService.CreateAccount(ctx, &account),
		http.StatusBadRequest); err != nil {
		return
	}
	httpbs.WriteJson(w, http.StatusCreated, account, "Account created succesfully")
}
