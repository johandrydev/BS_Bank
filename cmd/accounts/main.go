package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"BlueSoftBank/internal/database"
	"BlueSoftBank/internal/handlers"
	httpbs "BlueSoftBank/internal/pkg/http_bs"
)

func main() {
	// import the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	// create a new mux router
	mux := http.NewServeMux()

	db, err := database.CreateConnection()
	if err != nil {
		panic(err)
	}
	accountHandlers := handlers.NewAccountHandler(db)
	mux.HandleFunc(
		httpbs.CreateRoute(http.MethodPost, "api/account"), accountHandlers.CreateAccount,
	)

	// start the server
	port := os.Getenv("ACCOUNT_SERVICE_PORT")
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
