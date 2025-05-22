package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct{
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err !=nil{
		log.Printf("failed to write JSON response: %v", err)
	}

}

func writeError(w http.ResponseWriter, status int, err error){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := errorResponse{
		Error: err.Error(),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil{
		log.Printf("failed to write error response: %v", err)
	}
}