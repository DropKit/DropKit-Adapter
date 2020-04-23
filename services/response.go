package services

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
