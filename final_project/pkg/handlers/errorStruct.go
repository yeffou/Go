package handlers

import (
	"encoding/json"
	"final_project/pkg/models"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: message})
}
