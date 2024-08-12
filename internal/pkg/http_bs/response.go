package httpbs

import (
	"BlueSoftBank/internal/structures"
	"encoding/json"
	"log"
	"net/http"
)

func ValidateErr(w http.ResponseWriter, err error, status int) error {
	if err != nil {
		response := structures.HttpResponse{
			Message: err.Error(),
		}
		log.Println("request error", err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
	}
	return err
}

func WriteJson(w http.ResponseWriter, status int, data any, message string) {
	response := structures.HttpResponse{
		Data:    data,
		Message: message,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
