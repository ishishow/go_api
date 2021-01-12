package service

import (
	"encoding/json"
	"net/http"
)

type GachaDrawRequest struct {
	results []GachaResult
}

type GachaResult struct {
	CharacterID int
	Name        string
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	rep, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(rep))
	}
}
