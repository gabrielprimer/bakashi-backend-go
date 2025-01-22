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
	// Acesse o arquivo episodes.json a partir do diretório correto
	file, err := os.Open("data/episodes.json")
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

	// Define o tipo de conteúdo como JSON
	w.Header().Set("Content-Type", "application/json")
	// Envia a resposta como JSON
	json.NewEncoder(w).Encode(episodes)
}
