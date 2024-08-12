package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new mux router
	mux := http.NewServeMux()

	mux.HandleFunc(http.MethodGet+" api/movements", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World from movements")
	})

	// start the server
	if err := http.ListenAndServe(":8082", mux); err != nil {
		panic(err)
	}
}
