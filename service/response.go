package service

import (
	"encoding/json"
	"net/http"
)

type GachaDrawRequest struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	CharacterID int    `json:"character_id"`
	Name        string `json:"name"`
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
