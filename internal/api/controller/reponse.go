package controller

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	type respErr struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	resp := respErr{
		false,
		message,
	}
	RespondWithJSON(w, code, resp)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Write(response)
}
