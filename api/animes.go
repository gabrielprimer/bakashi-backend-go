package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"backend-bakashi-go/api/types"
)

// Handler para retornar os animes
func AnimesHandler(w http.ResponseWriter, r *http.Request) {
	// Acesse o arquivo animes.json a partir do diretório correto
	file, err := os.Open("data/animes.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir animes.json: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var animes types.Animes
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&animes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar animes.json: %v", err), http.StatusInternalServerError)
		return
	}

	// Define o tipo de conteúdo como JSON
	w.Header().Set("Content-Type", "application/json")
	// Envia a resposta como JSON
	json.NewEncoder(w).Encode(animes)
}
