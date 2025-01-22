package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"backend-bakashi-go/api/types"
)

// Handler para retornar os episódios
func EpisodesHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./data/episodes.json") // Ajustado para funcionar corretamente
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir episodes.json: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var episodes types.Episodes
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&episodes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar episodes.json: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}
