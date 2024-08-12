package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new mux router
	mux := http.NewServeMux()

	mux.HandleFunc(http.MethodGet+" api/clients", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World from clients")
	})

	// start the server
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
