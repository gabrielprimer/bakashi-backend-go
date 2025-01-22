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
	file, err := os.Open("./data/animes.json") // Ajustado para funcionar corretamente
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
}
