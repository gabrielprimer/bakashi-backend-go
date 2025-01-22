package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"

	"backend-bakashi-go/api/types"
)

// Handler para retornar os episódios
func EpisodesHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o diretório de trabalho atual
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter diretório de trabalho: %v", err), http.StatusInternalServerError)
		return
	}

	// Abre o arquivo JSON de episódios
	file, err := os.Open(fmt.Sprintf("%s/data/episodes.json", dir))
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

	// Define o tipo de conteúdo como JSON e envia a resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}
