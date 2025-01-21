package main

import (
	"encoding/json"
	"net/http"
	"os"
	"backend-bakashi-go/types"
)

// Função para carregar os episódios
func loadEpisodes() ([]types.Episode, error) {
	data, err := os.ReadFile("data/episodes.json")
	if err != nil {
		return nil, err
	}
	var episodes struct {
		Episodes []types.Episode `json:"episodes"`
	}
	err = json.Unmarshal(data, &episodes)
	if err != nil {
		return nil, err
	}
	return episodes.Episodes, nil
}

// Handler para listar os episódios
func EpisodesHandler(w http.ResponseWriter, r *http.Request) {
	episodes, err := loadEpisodes()
	if err != nil {
		http.Error(w, "Erro ao carregar os episódios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"episodes": episodes})
}
