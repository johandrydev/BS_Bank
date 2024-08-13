package main

import (
	"BlueSoftBank/internal/database"
	"BlueSoftBank/internal/handlers"
	httpbs "BlueSoftBank/internal/pkg/http_bs"
	"net/http"
)

func main() {
	// create a new mux router
	mux := http.NewServeMux()

	db, err := database.CreateConnection()
	if err != nil {
		panic(err)
	}
	accountHandlers := handlers.NewMovementHandler(db)
	mux.HandleFunc(
		httpbs.CreateRoute(http.MethodPost, "api/movement"), accountHandlers.CreateMovement,
	)

	// start the server
	if err := http.ListenAndServe(":8082", mux); err != nil {
		panic(err)
	}
}
