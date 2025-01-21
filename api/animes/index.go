package main

import (
	"encoding/json"
	"net/http"
	"os"
	"backend-bakashi-go/types"
)

// Função para carregar os animes
func loadAnimes() ([]types.Anime, error) {
	data, err := os.ReadFile("data/animes.json")
	if err != nil {
		return nil, err
	}
	var animes struct {
		Animes []types.Anime `json:"animes"`
	}
	err = json.Unmarshal(data, &animes)
	if err != nil {
		return nil, err
	}
	return animes.Animes, nil
}

// Handler para listar os animes
func AnimesHandler(w http.ResponseWriter, r *http.Request) {
	animes, err := loadAnimes()
	if err != nil {
		http.Error(w, "Erro ao carregar os animes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"animes": animes})
}
