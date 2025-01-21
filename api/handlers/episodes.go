package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/api/types"
)

func EpisodesHandler(w http.ResponseWriter, r *http.Request) {
	episodes, err := loadEpisodes()
	if err != nil {
		log.Printf("Erro ao carregar os episódios: %v", err)
		http.Error(w, "Erro ao carregar os episódios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"episodes": episodes})
}

func loadEpisodes() ([]types.Episode, error) {
	dataPath := "api/data/episodes.json"
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}
	var result struct {
		Episodes []types.Episode `json:"episodes"`
	}
	err = json.Unmarshal(data, &result)
	return result.Episodes, err
}
